using System;
using System.ComponentModel;
using System.Windows;
using AppLib;

namespace TesterApp
{
    /// <summary>
    ///     Interaction logic for Exit.xaml
    /// </summary>
    public partial class Exit : Window
    {
        private readonly TestAnswers testAnswers;
        private readonly Window parent;
        private bool flag;
        private string id;

        public Exit(Window parent, string id, TestAnswers testAnswers, bool areAllAnswersGot)
        {
            InitializeComponent();
            flag = true;
            this.testAnswers = testAnswers;
            this.parent = parent;
            this.id = id;
            Topmost = true;
            if (!areAllAnswersGot) TextBlock.Text = "Are you sure you want to exit?" +
                    "\nYou have not answered all the questions.";
            else TextBlock.Text = "Are you sure you want to exit?";
        }

        private void YesButton_Click(object sender, RoutedEventArgs e)
        {
            if (!(Server.SendData(id, testAnswers)))
            {
                MessageBox.Show("Lost connection to server. " +
                "\nAsk your teacher for help");
                MessageBox.Show("Testing was interrupted");
            }
            flag = false;
            parent.Close();
            Close();
        }

        private void Window_Deactivated(object sender, EventArgs e)
        {
            if (flag) Topmost = true;
            Topmost = false;
        }

        private void Window_Closing(object sender, CancelEventArgs e)
        {
            if (flag) e.Cancel = true;
        }

        private void NoButton_Click(object sender, RoutedEventArgs e)
        {
            flag = false;
            Close();
        }
    }
}