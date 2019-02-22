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
        public Button[] numbers;
        public Button exit;
        public Answers answers;

        public Window1(Student student)
        {
            InitializeComponent();
            this.Height = SystemParameters.FullPrimaryScreenHeight;
            this.Width = SystemParameters.FullPrimaryScreenWidth;
            this.WindowState = WindowState.Maximized;
            questions = Server.GetQuestions();
            answers = new Answers(questions.Length);
            ShowQuestion(0);
        }

        public void ShowQuestion(int number)
        {
            int num = 0;
            Question question = questions[number];
            foreach (Question q in questions)
                if (q.Section() == question.Section()) num++;
            numbers = new Button[num];
            for (int i = 0; i < num; i++)
            {
                numbers[i] = new Button();
                //numbers[i].Height = 50;
                numbers[i].Width = dockpanel1.Width / num;
                numbers[i].Content = "" + (i + 1);
                //numbers[i].Margin = new Thickness(numbers[i].Width*i,
                //this.Height - 50, 0, 0);
                dockpanel1.Children.Add(numbers[i]);
            }

            if (question.Section() == 3) ShowWriting(number);
            else if (question.Type() == 0) ShowType1(number);
            else ShowType2(number);
        }

        public void ShowWriting(int number)
        {

        }

        public void ShowType1(int number)
        {

        }

        public void ShowType2(int number)
        {

        }

    }
}
