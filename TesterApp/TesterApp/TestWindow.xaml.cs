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
            Textbox.BorderThickness = new Thickness(3);
        }

        private void ShowQuestion(int number)
        {
            Grid.Children.Clear();
            TestWindowDockpanel.Children.Clear();
            Reading.Background = Brushes.White;
            Listening.Background = Brushes.White;
            Writing.Background = Brushes.White;
            var num = 0;
            ActualNumber = number;
            var question = Questions[number];
            _actualSection = question.Section;
            var lightBlue = new Color
            {
                A = 1,
                B = 255,
                G = 230,
                R = 217
            };
            switch (question.Section)
            {
                case 1:
                    Listening.Background = new SolidColorBrush(lightBlue);
                    break;
                case 2:
                    Reading.Background = new SolidColorBrush(lightBlue);
                    break;
                case 3:
                    Writing.Background = new SolidColorBrush(lightBlue);
                    break;
                default:
                    throw new ArgumentOutOfRangeException();
            }

            foreach (var q in Questions)
                if (q.Section == question.Section)
                    num++;
            QButtons = new Button[num];
            for (var i = 0; i < num; i++)
            {
                QButtons[i] = new Button
                {
                    Name = "l" + i,
                    Width = TestWindowDockpanel.Width / num,
                    Content = "  " + (i + 1) + "  "
                };
                QButtons[i].Click += ButtonOnClick;
                TestWindowDockpanel.Children.Add(QButtons[i]);
            }

            if (question.Section == 3) ShowWriting();
            else ShowQuestion();
        }

        private void ShowWriting()
        {
            Textblock2.Text = "Введите ответ в поле ниже.";
            Textblock.Text = Questions[ActualNumber].Text;
            Textblock.Height = (Height - 70) / 3;
            Textbox = new TextBox
            {
                TextWrapping = TextWrapping.Wrap,
                VerticalScrollBarVisibility = ScrollBarVisibility.Visible,
                AcceptsReturn = true,
                Height = (Height - 70) / 3 * 2,
                Text = Answers[ActualNumber]
            };
            Grid.Children.Add(Textbox);
        }

        private void ShowQuestion()
        {
            Textblock2.Text = "Введите ответ в поле ниже. " +
                              "Если ответ подразумевает собой несколько вариантов ответов, введите их номера/буквы" +
                              " подряд без пробелов в том порядке, в каком они расположены в задании.";
            Textblock.Text = Questions[ActualNumber].Text;
            Textblock.Text = Questions[ActualNumber].Text;
            Textblock.Height = (Height - 70) / 2;
            Textbox = new TextBox
            {
                Height = (Height - 70) / 2,
                Text = Answers[ActualNumber]
            };
            Grid.Children.Add(Textbox);
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
            Answers[ActualNumber] = Textbox.Text;
        }
    }
}