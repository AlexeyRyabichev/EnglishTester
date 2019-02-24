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
    /// Логика взаимодействия для Exit.xaml
    /// </summary>
    public partial class Exit : Window
    {
        public Student student;
        public Window parent;
        private bool flag;
        public Exit(Student student, Window parent)
        {
            InitializeComponent();
            flag = true;
            this.student = student;
            this.parent = parent;
            this.Topmost = true;
        }

        private void button1_Click(object sender, RoutedEventArgs e)
        {
            Server.SendData(student);
            flag = false;
            parent.Close();
            this.Close();
        }

        private void Window_Deactivated(object sender, EventArgs e)
        {
            if (flag) this.Topmost = true;
        }

        private void Window_Closing(object sender, System.ComponentModel.CancelEventArgs e)
        {
            if (flag) e.Cancel = true;
        }

        private void button2_Click_1(object sender, RoutedEventArgs e)
        {
            flag = false;
            this.Close();
        }
    }
}
