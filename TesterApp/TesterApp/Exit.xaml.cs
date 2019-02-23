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
        Student student;
        Window parent;
        public Exit(Student student, Window parent)
        {
            InitializeComponent();
            this.student = student;
            this.parent = parent;
        }

        private void button1_Click(object sender, RoutedEventArgs e)
        {
            this.Close();
        }
        private void button2_Click(object sender, RoutedEventArgs e)
        {
            Server.SendData(student);
            parent.Close();
            this.Close();
        }
    }
}
