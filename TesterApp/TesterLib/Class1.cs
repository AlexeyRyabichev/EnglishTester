using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace TesterLib
{
    public class Student
    {
        private string email;
        private string password;
        private Answers answers;
        public Student(string email, string password)
        {
            if (!(email.Contains("@edu.hse.ru")) || (password.Length < 4))
                throw new ArgumentOutOfRangeException("Неверно введённые данные");
            if (!(Authentication(email, password)))
                throw new ArgumentOutOfRangeException("Неверно введённые данные");
            this.email = email;
            this.password = password;
        }
        public void AddAnswers(Answers answers)
        {
            this.answers = answers;
        } 
        public bool Authentication(string email, string password)
        {
            return true;
        }
    }

    public class Answers
    {
        private string[] answers;
        public Answers(int size)
        {
            if (size <= 0) throw new IndexOutOfRangeException();
            answers = new string[size];
        }
        public void AddAnswer(int number, string answer)
        {
            answers[number] = answer;
        }
    }

    public class Question
    {
        private int type;
        private int section;
        private string text;
        private string[] answers;

        public Question(string question, int type, int section, string[] answers)
        {
            this.text = question;
            this.type = type;
            this.answers = answers;
            this.section = section;
        }
        public int Type()
        {
            return type;
        }
        public string Text()
        {
            return text;
        }
        public int Section()
        {
            return section;
        }
        public string[] Answers()
        {
            return answers;
        }
    }
}
