﻿using System;
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
        private readonly Student student;
        private readonly Window parent;
        private bool flag;

        public Exit(Student student, Window parent)
        {
            InitializeComponent();
            flag = true;
            this.student = student;
            this.parent = parent;
            Topmost = true;
        }

        private void YesButton_Click(object sender, RoutedEventArgs e)
        {
            Server.SendData(student);
            flag = false;
            parent.Close();
            Close();
        }

        private void Window_Deactivated(object sender, EventArgs e)
        {
            if (flag) Topmost = true;
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