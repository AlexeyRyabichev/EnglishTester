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
using System.Windows.Navigation;
using System.Windows.Shapes;
using TesterLib;
using ServerLib;

namespace TesterApp
{
    /// <summary>
    /// Логика взаимодействия для MainWindow.xaml
    /// </summary>
    public partial class MainWindow : Window
    {
        public Student student;

        public MainWindow()
        {
            InitializeComponent();
            this.WindowState = WindowState.Maximized;
            button1.Click += TryToStart;
        }
        void TryToStart(object sender, RoutedEventArgs e)
        {
            string email = textbox1.Text;
            string password = textbox2.Text;
            try
            {
                string id = Server.Authentication(email, password);
                if (id == "")
                    throw new ArgumentOutOfRangeException("Неверно введённые данные");
                student = new Student(email, password, id);
                Window1 testerWindow = new Window1(student);
                testerWindow.Show();
                this.Close();
            }
            catch (Exception ex)
            {
                MessageBox.Show(ex.Message);
            }
        }

    }
}
