using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.Windows;
using System.Windows.Controls;
using System.Windows.Data;
using System.Windows.Documents;
using System.Windows.Input;
using System.Windows.Media;
using System.Windows.Media.Imaging;
using System.Windows.Shapes;
using TesterLib;
using ServerLib;

namespace TesterApp
{
    /// <summary>
    /// Логика взаимодействия для Window1.xaml
    /// </summary>
    public partial class Window1 : Window
    {
        public Question[] questions;
        public Button[] q_buttons;
        public Button exit;
        public TextBox textbox;
        public CheckBox[] checkboxes;
        public Answers[] answers;
        public Student student;
        public int actual_number;

        public Window1(Student student)
        {
            InitializeComponent();
            this.Height = SystemParameters.FullPrimaryScreenHeight;
            this.Width = SystemParameters.FullPrimaryScreenWidth;
            this.WindowState = WindowState.Maximized;
            this.BorderThickness = new Thickness(0);
            questions = Server.GetQuestions();
            answers = new Answers[questions.Length];
            this.student = student;
            ShowQuestion(0);
            textbox.BorderThickness = new Thickness(3);
        }

        public void ShowQuestion(int number)
        {
            grid.Children.Clear();
            dockpanel1.Children.Clear();
            reading.Background = Brushes.White;
            listening.Background = Brushes.White;
            writing.Background = Brushes.White;
            int num = 0;
            actual_number = number;
            Question question = questions[number];
            if (answers[number] == null)
            {
                int count;
                if (question.Type == 0) count = 1;
                else count = question.Type;
                answers[number] = new Answers(count);
            }
            switch(question.Section)
            {
                case 1:
                    listening.Background = Brushes.Aquamarine;
                    break;
                case 2:
                    reading.Background = Brushes.Aquamarine;
                    break;
                case 3:
                    writing.Background = Brushes.Aquamarine;
                    break;
                default:
                    throw new ArgumentOutOfRangeException();
            }
            foreach (Question q in questions)
                if (q.Section == question.Section) num++;
            q_buttons = new Button[num];
            for (int i = 0; i < num; i++)
            {
                q_buttons[i] = new Button();
                //numbers[i].Height = 50;
                q_buttons[i].Width = dockpanel1.Width / num;
                q_buttons[i].Content = "  " + (i + 1) + "  ";
                //numbers[i].Margin = new Thickness(numbers[i].Width*i,
                //this.Height - 50, 0, 0);
                dockpanel1.Children.Add(q_buttons[i]);
            }

            if (question.Section == 3) ShowWriting();
            else if (question.Type == 0) ShowType1();
            else ShowType2();
        }

        public void ShowWriting()
        {
            textblock.Text = questions[actual_number].Text;
            textblock.Height = (this.Height - 70) / 3;
            TextBox textbox = new TextBox();
            textbox.Height = (this.Height - 70) / 3 * 2;
            grid.Children.Add(textbox);
        }

        public void ShowType1()
        {
            textblock.Text = questions[actual_number].Text;
            textblock.Text = questions[actual_number].Text;
            textblock.Height = (this.Height - 70) / 2;
            textbox = new TextBox();
            textbox.Height = (this.Height - 70) / 2;
            grid.Children.Add(textbox);
        }

        public void ShowType2()
        {

            textblock.Text = questions[actual_number].Text;
            checkboxes = new CheckBox[questions[actual_number].Type];
            Answers ans = questions[actual_number].Answers;
            for (int i = 0; i < checkboxes.Length; i++)
            {
                //checkboxes[i].Content = ans[i];
            }

        }

        private void TextBox_TextInput(object sender, KeyEventArgs e)
        {
            answers[actual_number].AddAnswer(textbox.Text);
        }

        private void checkBox_Checked(object sender, RoutedEventArgs e)
        {
            answers[actual_number].AddAnswer((string)((CheckBox)sender).Content);
        }

        private void checkBox_Unchecked(object sender, RoutedEventArgs e)
        {
            answers[actual_number].DeleteAnswer((string)((CheckBox)sender).Content);
        }

        private void reading_Click(object sender, RoutedEventArgs e)
        {
            int i;
            for (i = 0; i < questions.Length; i++)
                if (questions[i].Section == 2)
                {
                    ShowQuestion(i);
                    break;
                }
        }

        private void listening_Click(object sender, RoutedEventArgs e)
        {
            int i;
            for (i = 0; i < questions.Length; i++)
                if (questions[i].Section == 1)
                {
                    ShowQuestion(i);
                    break;
                }
        }

        private void writing_Click(object sender, RoutedEventArgs e)
        {
            int i;
            for (i = 0; i < questions.Length; i++)
                if (questions[i].Section == 3) {
                    ShowQuestion(i);
                    break;
                }
        }

        private void submit_Click(object sender, RoutedEventArgs e)
        {
            student.AddAnswers(answers);
            Exit exit = new Exit(student, this);
        }
    }
}
