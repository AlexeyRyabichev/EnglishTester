﻿using System;
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
        private bool flag;
        private Brush defaultColor;
        private bool areAllAnswersGot;
        private int actualNumber;
        private int question_count;
        private string[] answers;
        public Button[] QuestionButtons;
        private Student student;
        public TextBox TextBox;
        public RadioButton[] RadioButtons;
        public StackPanel RadioPanel;
        public Test Test;
        public TestAnswers testAnswers;

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
            Test = Server.GetTest(student.ID);
            if (Test == null) throw new DivideByZeroException();
            question_count = Test.Reading.Questions.Length 
                + Test.BaseQuestions.Length + 1;
            answers = 
                new string[question_count];
            this.student = student;
            areAllAnswersGot = false;
            RadioPanel = new StackPanel();
            AddButtons();
            ShowQuestion(0);
        }

        private void ShowQuestion(int number)
        {
            AnswerPanel.Children.Clear();
            TextBox = null;
            actualNumber = number;
            if (actualNumber < Test.BaseQuestions.Length)
            {
                ShowQuestion_Base();
            }
            else if (actualNumber < question_count - 1)
            {
                ShowQuestion_Reading();
            }
            else
            {
                ShowQuestion_Writing();
            }
        }

        private void ShowQuestion_Writing()
        {
            Textblock.Text = "Writing.\n\n";
            Textblock.Text += Test.Writing;
            gridText.Height = (Height - 20) / 5 * 2;
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
            Textblock.Text = "Reading. Excercize " + (actualNumber - Test.BaseQuestions.Length + 1) + ".\n\n";
            Textblock.Text += Test.Reading.Question;
            Textblock.Text += "\n\n" + 
                Test.Reading.Questions[actualNumber - Test.BaseQuestions.Length].Question;
            gridText.Height = (Height - 20) / 5 * 4;
            Textblock2.Text = "Choose the correct answer:";
            RadioButtons = new RadioButton[4];
            RadioPanel.Children.Clear();
            RadioButtons[0] = new RadioButton
            {
                Content = "" + 
                Test.Reading.Questions[actualNumber - Test.BaseQuestions.Length].optionA,
                Name = "A1"
            };
            RadioButtons[1] = new RadioButton
            {
                Content = "" +
                Test.Reading.Questions[actualNumber - Test.BaseQuestions.Length].optionB,
                Name = "B2"
            };
            RadioButtons[2] = new RadioButton
            {
                Content = "" +
                Test.Reading.Questions[actualNumber - Test.BaseQuestions.Length].optionC,
                Name = "C3"
            };
            RadioButtons[3] = new RadioButton
            {
                Content = "" +
                Test.Reading.Questions[actualNumber - Test.BaseQuestions.Length].optionD,
                Name = "D4"
            };
            for (int i = 0; i < 4; i++)
            {
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

        private void ShowQuestion_Base()
        {
            Textblock.Text = "Excercize " + (actualNumber + 1) + ".\n\n";
            Textblock.Text += Test.BaseQuestions[actualNumber].Question;           
            gridText.Height = (Height - 20) / 5 * 4;          
            Textblock2.Text = "Choose the correct answer:";
            RadioButtons = new RadioButton[4];
            RadioPanel.Children.Clear();
            RadioButtons[0] = new RadioButton
            {
                Content = "" +
                Test.BaseQuestions[actualNumber].optionA,
                Name = "A1"
            };
            RadioButtons[1] = new RadioButton
            {
                Content = "" +
                Test.BaseQuestions[actualNumber].optionB,
                Name = "B2"
            };
            RadioButtons[2] = new RadioButton
            {
                Content = "" +
                Test.BaseQuestions[actualNumber].optionC,
                Name = "C3"
            };
            RadioButtons[3] = new RadioButton
            {
                Content = "" +
                Test.BaseQuestions[actualNumber].optionD,
                Name = "D4"
            };
            for (int i = 0; i < 4; i++)
            {
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

        private void Base_Click(object sender, RoutedEventArgs e)
        {
            WriteAnswers();
            ShowQuestion(0);
        }

        private void Reading_Click(object sender, RoutedEventArgs e)
        {
            WriteAnswers();
            ShowQuestion(Test.BaseQuestions.Length);
        }


        private void Writing_Click(object sender, RoutedEventArgs e)
        {
            WriteAnswers();
            ShowQuestion(question_count - 1);
        }

        private void Submit_Click(object sender, RoutedEventArgs e)
        {
            WriteAnswers();
            student.AddAnswers(answers);
            flag = false;
            FormateAnswers();
            var exit = new Exit(this, student.ID, testAnswers, areAllAnswersGot);
            exit.ShowDialog();
        }

        private void Window_Closing(object sender, CancelEventArgs e)
        {
            if (flag) e.Cancel = true;
        }

        private void Window_Deactivated(object sender, EventArgs e)
        {
            if (flag) Topmost = true;
            //Topmost = false;
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
           
            QuestionButtons = new Button[question_count];
            for (var i = 0; i < Test.BaseQuestions.Length; i++)
            {
                    QuestionButtons[i] = new Button
                    {
                        Name = "q" + i,
                        Content = "" + (i + 1) + "",
                        Margin = new Thickness(5),
                        MinWidth = 25,
                        MaxWidth = Height
                    };
                    QuestionButtons[i].Click += ButtonOnClick;

                    BasePanel.Children.Add(QuestionButtons[i]);
            }
            for (var i = Test.BaseQuestions.Length; i < question_count - 1; i++)
            {
                QuestionButtons[i] = new Button
                {
                    Name = "q" + i,
                    Content = "" + (i - Test.BaseQuestions.Length + 1) + "",
                    Margin = new Thickness(5),
                    MinWidth = 25,
                    MaxWidth = Height
                };
                QuestionButtons[i].Click += ButtonOnClick;

                ReadingPanel.Children.Add(QuestionButtons[i]);
            }

            QuestionButtons[question_count - 1] = new Button
            {
                Name = "q" + (question_count - 1),
                Content = "1",
                Margin = new Thickness(5),
                MinWidth = 35,
                MaxWidth = 45
            };
            QuestionButtons[question_count - 1].Click += ButtonOnClick;
            WritingPanel.Children.Add(QuestionButtons[question_count - 1]);

        }

        private void CheckAnswers()
        {
            int num = answers.Length;
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

       
        private void FormateAnswers()
        {
            testAnswers = new TestAnswers();
            testAnswers.Base = new VariableAnswer[Test.BaseQuestions.Length];
            for (int i = 0; i < testAnswers.Base.Length; i++)
            {
                testAnswers.Base[i] = new VariableAnswer(Test.BaseQuestions[i].Id,
                    FromIntToABCD(answers[i]));
            }
            testAnswers.Reading = new VariableAnswer[Test.Reading.Questions.Length];
            for (int i = 0; i < testAnswers.Reading.Length; i++)
            {
                testAnswers.Reading[i] = new VariableAnswer(Test.Reading.Questions[i].Id,
                    FromIntToABCD(answers[Test.BaseQuestions.Length + i]));
            }
            testAnswers.Writing = answers[answers.Length - 1];
        }

        private string FromIntToABCD(string str)
        {
            switch (str)
            {
                case "1":
                    return "A";
                case "2":
                    return "B";
                case "3":
                    return "C";
                case "4":
                    return "D";
                default:
                    return "";
            }
        }
    }
}