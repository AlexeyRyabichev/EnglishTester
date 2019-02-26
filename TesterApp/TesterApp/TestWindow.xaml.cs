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
        private Brush defaultColor;
        private bool areAllAnswersGot;
        private int actualNumber;
        private string[] answers;
        public Button[] QuestionButtons;
        private Question[] questions;
        private Student student;
        public TextBox TextBox;
        public RadioButton[] RadioButtons;
        public StackPanel RadioPanel;

        public TestWindow(Student student)
        {
            InitializeComponent();
            Height = SystemParameters.FullPrimaryScreenHeight;
            Width = SystemParameters.FullPrimaryScreenWidth;
            WindowState = WindowState.Maximized;
            Topmost = true;
            BorderThickness = new Thickness(0);
            flag = true;
            defaultColor = Reading.Background;
            questions = Server.GetQuestions();
            answers = new string[questions.Length];
            this.student = student;
            areAllAnswersGot = false;
            RadioPanel = new StackPanel();
            AddButtons();
            ShowQuestion(0);
        }

        private void ShowQuestion(int number)
        {
            AnswerPanel.Children.Clear();
            actualNumber = number;
            var question = questions[number];
            actualSection = question.Section;
            if (actualSection == 2) ShowQuestion_Reading();
            else ShowQuestion_Writing();
        }

        private void ShowQuestion_Writing()
        {
            Textblock.Text = questions[actualNumber].Text;
            Textblock.Height = (Height - 20) / 3;
            Textblock2.Text = "Type your answer in the box below:";
            TextBox = new TextBox
            {
                TextWrapping = TextWrapping.Wrap,
                VerticalScrollBarVisibility = ScrollBarVisibility.Visible,
                AcceptsReturn = true,
                Text = answers[actualNumber],
                Margin = new Thickness(5),
                BorderThickness = new Thickness(2),
                VerticalContentAlignment = VerticalAlignment.Top,
            };
            AnswerPanel.Children.Add(TextBox);
        }

        private void ShowQuestion_Reading()
        {
            Textblock.Text = questions[actualNumber].Text;
            Textblock.Height = (Height - 20) / 5 * 4;
            Textblock2.Text = "Choose the correct answer:";
            RadioButtons = new RadioButton[4];
            RadioPanel.Children.Clear();
            for (int i = 0; i < 4; i++)
            {
                RadioButtons[i] = new RadioButton
                {
                    Content = "" + (i + 1),
                    Name = "r" + i
                };
                RadioButtons[i].Checked += RadioButtonOnClick;
                RadioPanel.Children.Add(RadioButtons[i]);
            }
            AnswerPanel.Children.Add(RadioPanel);
            if ((answers[actualNumber] != "") && (answers[actualNumber] != null))
            {
                int.TryParse(answers[actualNumber], out int num);
                RadioButtons[num].IsChecked = true;
            }
        }


        private void Reading_Click(object sender, RoutedEventArgs e)
        {
            WriteAnswers();
            CheckAnswers();
            int i;
            for (i = 0; i < questions.Length; i++)
                if (questions[i].Section == 2)
                {
                    ShowQuestion(i);
                    break;
                }
        }


        private void Writing_Click(object sender, RoutedEventArgs e)
        {
            WriteAnswers();
            CheckAnswers();
            int i;
            for (i = 0; i < questions.Length; i++)
                if (questions[i].Section == 3)
                {
                    ShowQuestion(i);
                    break;
                }
        }

        private void Submit_Click(object sender, RoutedEventArgs e)
        {
            WriteAnswers();
            CheckAnswers();
            student.AddAnswers(answers);
            flag = false;
            var exit = new Exit(this, student, areAllAnswersGot);
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
            var button = (Button) sender;
            int.TryParse(button.Name.Substring(1), out var index);
            ShowQuestion(index);
        }

        private void RadioButtonOnClick(object sender, EventArgs eventArgs)
        {
            RadioButton radioButton = (RadioButton)sender;
            string stringNumber = radioButton.Name.Substring(1);
            int.TryParse(stringNumber, out int number);
            answers[actualNumber] = "" + number;
        }

        private void WriteAnswers()
        {
            if (TextBox != null) answers[actualNumber] = TextBox.Text;
            CheckAnswers();
        }

        private void AddButtons()
        {
            int num = questions.Length;
            QuestionButtons = new Button[num];
            int readingCount = 0, writingCount = 0;
            for (var i = 0; i < num; i++)
            {
                if (questions[i].Section == 2)
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

        private void CheckAnswers()
        {
            int num = questions.Length;
            areAllAnswersGot = true;
            for (var i = 0; i < num; i++)
            {
                if ((answers[i] != null) && (answers[i] != ""))
                    QuestionButtons[i].Background = Brushes.LightSeaGreen;
                else
                {
                    QuestionButtons[i].Background = defaultColor;
                    areAllAnswersGot = false;
                }
            }
        }
    }
}