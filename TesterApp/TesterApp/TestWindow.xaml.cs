using System;
using System.ComponentModel;
using System.Windows;
using System.Windows.Controls;
using System.Windows.Media;
using AppLib;

namespace TesterApp
{
    /// <summary>
    ///     Interaction logic for TestWindow.xaml
    /// </summary>
    public partial class TestWindow : Window
    {
        private int _actualSection;
        private bool _flag;
        public int ActualNumber;
        public string[] Answers;
        public Button exit;
        public Button[] QButtons;
        public Question[] Questions;
        public Student Student;
        public TextBox TextBox;
        public RadioButton[] RadioButtons;

        public TestWindow(Student student)
        {
            InitializeComponent();
            Height = SystemParameters.FullPrimaryScreenHeight;
            Width = SystemParameters.FullPrimaryScreenWidth;
            WindowState = WindowState.Maximized;
            Topmost = true;
            BorderThickness = new Thickness(0);
            _flag = true;
            Questions = Server.GetQuestions();
            Answers = new string[Questions.Length];
            Student = student;
            ShowQuestion(0);
            TextBox.BorderThickness = new Thickness(3);
        }

        private void ShowQuestion(int number)
        {
            Grid.Children.Clear();
            TestWindowDockpanel.Children.Clear();
            var num = 0;
            ActualNumber = number;
            var question = Questions[number];
            _actualSection = question.Section;

            switch (question.Section)
            {
                case 1:
                    Listening.Background = Brushes.DarkBlue;
                    break;
                case 2:
                    Reading.Background = Brushes.DarkBlue;
                    break;
                case 3:
                    Writing.Background = Brushes.DarkBlue;
                    break;
                default:
                    throw new ArgumentOutOfRangeException();
            }

            AddButtons();
        }

        private void ShowWriting()
        {
            Textblock.Text = Questions[ActualNumber].Text;
            Textblock.Height = (Height - 70) / 3;
            TextBox = new TextBox
            {
                TextWrapping = TextWrapping.Wrap,
                VerticalScrollBarVisibility = ScrollBarVisibility.Visible,
                AcceptsReturn = true,
                Height = (Height - 70) / 3 * 2,
                Text = Answers[ActualNumber],
                Margin = new Thickness(5)
            };
            Grid.Children.Add(TextBox);
        }

        private void ShowQuestion()
        {
            Textblock.Text = Questions[ActualNumber].Text;
            Textblock.Text = Questions[ActualNumber].Text;
            Textblock.Height = (Height - 70) / 2;
            TextBox = new TextBox
            {
                Height = (Height - 70) / 2,
                Text = Answers[ActualNumber],
                Margin = new Thickness(5)
            };
            Grid.Children.Add(TextBox);
        }


        private void Reading_Click(object sender, RoutedEventArgs e)
        {
            Write();
            int i;
            for (i = 0; i < Questions.Length; i++)
                if (Questions[i].Section == 2)
                {
                    ShowQuestion(i);
                    break;
                }
        }

        private void Listening_Click(object sender, RoutedEventArgs e)
        {
            Write();
            int i;
            for (i = 0; i < Questions.Length; i++)
                if (Questions[i].Section == 1)
                {
                    ShowQuestion(i);
                    break;
                }
        }

        private void Writing_Click(object sender, RoutedEventArgs e)
        {
            Write();
            int i;
            for (i = 0; i < Questions.Length; i++)
                if (Questions[i].Section == 3)
                {
                    ShowQuestion(i);
                    break;
                }
        }

        private void Submit_Click(object sender, RoutedEventArgs e)
        {
            Write();
            Student.AddAnswers(Answers);
            _flag = false;
            var exit = new Exit(Student, this);
            exit.ShowDialog();
        }

        private void Window_Closing(object sender, CancelEventArgs e)
        {
            if (_flag) e.Cancel = true;
        }

        private void Window_Deactivated(object sender, EventArgs e)
        {
            if (_flag) Topmost = true;
        }


        private void ButtonOnClick(object sender, EventArgs eventArgs)
        {
            Write();
            var index = 0;
            var button = (Button) sender;
            int.TryParse(button.Name.Substring(1), out var number);
            while (Questions[index].Section != _actualSection || index != number)
            {
                index++;
                if (index >= Questions.Length) break;
            }

            if (index < Questions.Length) ShowQuestion(index);
        }

        private void Write()
        {
            Answers[ActualNumber] = TextBox.Text;
        }

        private void AddButtons()
        {
            int num = Questions.Length;
            Question question = Questions[0];
            QButtons = new Button[num];
            int re = 0, wr = 0, li = 0;
            for (var i = 0; i < num; i++)
            {
                if (Questions[i].Section == 1)
                {
                    QButtons[i] = new Button
                    {
                        Name = "q" + i,
                        Width = TestWindowDockpanel.Width / num,
                        Content = "  " + (li + 1) + "  ",
                        Margin = new Thickness(5),
                        MaxWidth = Height
                    };
                    li++;
                    QButtons[i].Click += ButtonOnClick;

                    ListeningPanel.Children.Add(QButtons[i]);
                }
                else if (Questions[i].Section == 2)
                {
                    QButtons[i] = new Button
                    {
                        Name = "q" + i,
                        Width = TestWindowDockpanel.Width / num,
                        Content = "  " + (re + 1) + "  ",
                        Margin = new Thickness(5),
                        MaxWidth = Height
                    };
                    re++;
                    QButtons[i].Click += ButtonOnClick;

                    ReadingPanel.Children.Add(QButtons[i]);
                }
                else
                {
                    QButtons[i] = new Button
                    {
                        Name = "q" + i,
                        Width = TestWindowDockpanel.Width / num,
                        Content = "  " + (wr + 1) + "  ",
                        Margin = new Thickness(5),
                        MaxWidth = Height
                    };
                    wr++;
                    QButtons[i].Click += ButtonOnClick;

                    WritingPanel.Children.Add(QButtons[i]);
                }
            }

            if (question.Section == 3) ShowWriting();
            else ShowQuestion();
        }
    }
}