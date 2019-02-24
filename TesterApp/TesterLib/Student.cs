using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using Newtonsoft.Json;

namespace TesterLib
{
    public class Student
    {
        public string Email { get; private set; }
        public string Password { get; private set; }
        public string[] Answers { get; private set; }
        public string ID { get; set; }
        private string id;
        public Student(string email, string password, string id)
        {
            Email = email;
            Password = password;
            ID = id;
        }
        public void AddAnswers(string[] answers)
        {
            this.Answers = new string[answers.Length];
            int i = 0;
            foreach (string s in answers)
            {
                if (s == null) this.Answers[i] = "";
                else this.Answers[i] = s;
                i++;
            }
        }
    }

    

    public class Question
    {
        //раздел. 1 - listening, 2 - reading, 3 - writing
        public string Text { get; private set; }
        //текст вопроса
        public int Section { get; private set; }


        public Question(string question, int section)
        {
            if (question == "") throw new ArgumentNullException();
            if ((section > 3) || (section < 0)) throw new ArgumentOutOfRangeException();          
            Text = question;
            Section = section;
        }

    }
}
