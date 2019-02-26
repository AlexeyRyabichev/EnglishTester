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
        private int actualSection;
        private bool flag;
        public int ActualNumber;
        public string[] Answers;
        public Button[] QuestionButtons;
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
            flag = true;
            Questions = Server.GetQuestions();
            Answers = new string[Questions.Length];
            Student = student;
            AddButtons();
            ShowQuestion(0);
        }

        private void ShowQuestion(int number)
        {
            Grid.Children.Clear();
            ActualNumber = number;
            var question = Questions[number];
            actualSection = question.Section;
            if (actualSection == 2) ShowQuestion_Reading();
            else ShowQuestion_Writing();
        }

        private void ShowQuestion_Writing()
        {
            Textblock.Text = Questions[ActualNumber].Text;
            Textblock.Height = (Height - 20) / 3;
            Textblock2.Text = "Type your answer in the box below:";
            TextBox = new TextBox
            {
                TextWrapping = TextWrapping.Wrap,
                VerticalScrollBarVisibility = ScrollBarVisibility.Visible,
                AcceptsReturn = true,
                Text = Answers[ActualNumber],
                Margin = new Thickness(5),
                BorderThickness = new Thickness(2),
                VerticalContentAlignment = VerticalAlignment.Top
            };
            Grid.Children.Add(TextBox);
        }

        private void ShowQuestion_Reading()
        {
            Textblock.Text = Questions[ActualNumber].Text;
            Textblock.Height = (Height - 20) / 5 * 4;
            Textblock2.Text = "Choose the correct answer:";
            TextBox = new TextBox
            {
                TextWrapping = TextWrapping.Wrap,
                VerticalScrollBarVisibility = ScrollBarVisibility.Visible,
                AcceptsReturn = true,
                Text = Answers[ActualNumber],
                Margin = new Thickness(5),
                BorderThickness = new Thickness(2),
                VerticalContentAlignment = VerticalAlignment.Top
            };
            Grid.Children.Add(TextBox);
        }


        private void Reading_Click(object sender, RoutedEventArgs e)
        {
            WriteAnswers();
            int i;
            for (i = 0; i < Questions.Length; i++)
                if (Questions[i].Section == 2)
                {
                    ShowQuestion(i);
                    break;
                }
        }


        private void Writing_Click(object sender, RoutedEventArgs e)
        {
            WriteAnswers();
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
            WriteAnswers();
            Student.AddAnswers(Answers);
            flag = false;
            var exit = new Exit(Student, this);
            exit.ShowDialog();
        }

        private void Window_Closing(object sender, CancelEventArgs e)
        {
            if (flag) e.Cancel = true;
        }

        private void Window_Deactivated(object sender, EventArgs e)
        {
            if (flag) Topmost = true;
        }


        private void ButtonOnClick(object sender, EventArgs eventArgs)
        {
            WriteAnswers();
            var index = 0;
            var button = (Button) sender;
            int.TryParse(button.Name.Substring(1), out var number);
            while (Questions[index].Section != actualSection || index != number)
            {
                index++;
                if (index >= Questions.Length) break;
            }

            if (index < Questions.Length) ShowQuestion(index);
        }

        private void WriteAnswers()
        {
            if (TextBox != null) Answers[ActualNumber] = TextBox.Text;
        }

        private void AddButtons()
        {
            int num = Questions.Length;
            QuestionButtons = new Button[num];
            int readingCount = 0, writingCount = 0;
            for (var i = 0; i < num; i++)
            {
                if (Questions[i].Section == 2)
                {
                    QuestionButtons[i] = new Button
                    {
                        Name = "q" + i,
                        Content = "  " + (readingCount + 1) + "  ",
                        Margin = new Thickness(5),
                        MaxWidth = Height
                    };
                    readingCount++;
                    QuestionButtons[i].Click += ButtonOnClick;

                    ReadingPanel.Children.Add(QuestionButtons[i]);
                }
                else
                {
                    QuestionButtons[i] = new Button
                    {
                        Name = "q" + i,
                        Content = "  " + (writingCount + 1) + "  ",
                        Margin = new Thickness(5),
                        MaxWidth = Height
                    };
                    writingCount++;
                    QuestionButtons[i].Click += ButtonOnClick;

                    WritingPanel.Children.Add(QuestionButtons[i]);
                }
            }
        }
    }
}