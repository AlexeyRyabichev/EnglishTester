using System;
using System.Windows;
using AppLib;

namespace TesterApp
{
    /// <summary>
    ///     Interaction logic for MainWindow.xaml
    /// </summary>
    public partial class MainWindow : Window
    {
        private Student _student;

        public MainWindow()
        {
            InitializeComponent();
            WindowState = WindowState.Maximized;
            LoginButton.Click += TryToStart;
        }

        private void TryToStart(object sender, RoutedEventArgs e)
        {
            var email = EmailTextBox.Text;
            var password = PasswordTextBox.Password;
            try
            {
                var id = Server.Authentication(email, password);
                if (id == "")
                    throw new FieldAccessException();
                _student = new Student(email, password, id);
                var testerWindow = new TestWindow(_student);
                testerWindow.Show();
                Close();
            }
            catch (FieldAccessException)
            {
                ErrorLabel.Content = "Wrong email or password";
            }
            catch (Exception ex)
            {
                MessageBox.Show(ex.ToString());
            }
        }
    }
}