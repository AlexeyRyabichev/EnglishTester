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
        public string[] answers;
        public Student student;
        public int actual_number;
        private bool flag;
        private int actual_section;

        public Window1(Student student)
        {
            InitializeComponent();
            this.Height = SystemParameters.FullPrimaryScreenHeight;
            this.Width = SystemParameters.FullPrimaryScreenWidth;
            this.WindowState = WindowState.Maximized;
            this.Topmost = true;
            this.BorderThickness = new Thickness(0);
            flag = true;
            questions = Server.GetQuestions();
            answers = new string[questions.Length];
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
            actual_section = question.Section;
            switch (question.Section)
            {
                case 1:
                    listening.Background = Brushes.LightSteelBlue;
                    break;
                case 2:
                    reading.Background = Brushes.LightSteelBlue;
                    break;
                case 3:
                    writing.Background = Brushes.LightSteelBlue;
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
                q_buttons[i].Name = "l" + i;
                q_buttons[i].Width = dockpanel1.Width / num;
                q_buttons[i].Content = "  " + (i + 1) + "  ";
                q_buttons[i].Click += ButtonOnClick;
                dockpanel1.Children.Add(q_buttons[i]);
            }

            if (question.Section == 3) ShowWriting();
            else ShowQuestion();
        }

        public void ShowWriting()
        {
            textblock2.Text = "Введите ответ в поле ниже.";
            textblock.Text = questions[actual_number].Text;
            textblock.Height = (this.Height - 70) / 3;
            textbox = new TextBox();
            textbox.TextWrapping = TextWrapping.Wrap;
            textbox.VerticalScrollBarVisibility = ScrollBarVisibility.Visible;
            textbox.AcceptsReturn = true;
            textbox.Height = (this.Height - 70) / 3 * 2;
            textbox.Text = answers[actual_number];
            grid.Children.Add(textbox);
        }

        public void ShowQuestion()
        {
            textblock2.Text = "Введите ответ в поле ниже. " +
                "Если ответ подразумевает собой несколько вариантов ответов, введите их номера/буквы" +
                " подряд без пробелов в том порядке, в каком они расположены в задании.";
            textblock.Text = questions[actual_number].Text;
            textblock.Text = questions[actual_number].Text;
            textblock.Height = (this.Height - 70) / 2;
            textbox = new TextBox();
            textbox.Height = (this.Height - 70) / 2;
            textbox.Text = answers[actual_number];
            grid.Children.Add(textbox);
        }


        private void reading_Click(object sender, RoutedEventArgs e)
        {
            Write();
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
            Write();
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
            Write();
            int i;
            for (i = 0; i < questions.Length; i++)
                if (questions[i].Section == 3) {
                    ShowQuestion(i);
                    break;
                }
        }

        private void submit_Click(object sender, RoutedEventArgs e)
        {
            Write();
            student.AddAnswers(answers);
            flag = false;
            Exit exit = new Exit(student, this);
            exit.ShowDialog();
        }

        private void Window_Closing(object sender, System.ComponentModel.CancelEventArgs e)
        {
            if (flag) e.Cancel = true;
        }

        private void Window_Deactivated(object sender, EventArgs e)
        {
            if (flag) this.Topmost = true;
        }


        private void ButtonOnClick(object sender, EventArgs eventArgs)
        {
            Write();
            int index = 0;
            var button = (Button)sender;
            int number;
            int.TryParse((button.Name).Substring(1), out number);
            while ((questions[index].Section != actual_section) || (index != number))
            {
                index++;
                if (index >= questions.Length) break;
            }
            if (index<questions.Length) ShowQuestion(index);
        }

        private void Write()
        {
            answers[actual_number] = textbox.Text;
        }
    }
}
