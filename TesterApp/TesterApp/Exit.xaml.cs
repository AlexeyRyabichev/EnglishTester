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
        private readonly Student _student;
        private readonly Window parent;
        private bool _flag;

        public Exit(Student student, Window parent)
        {
            InitializeComponent();
            _flag = true;
            _student = student;
            this.parent = parent;
            Topmost = true;
        }

        private void YesButton_Click(object sender, RoutedEventArgs e)
        {
            Server.SendData(_student);
            _flag = false;
            parent.Close();
            Close();
        }

        private void Window_Deactivated(object sender, EventArgs e)
        {
            if (_flag) Topmost = true;
        }

        private void Window_Closing(object sender, CancelEventArgs e)
        {
            if (_flag) e.Cancel = true;
        }

        private void NoButton_Click(object sender, RoutedEventArgs e)
        {
            _flag = false;
            Close();
        }
    }
}