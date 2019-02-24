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
        private string[] answers;
        private string id;
        public Student(string email, string password)
        {
            this.email = email;
            this.password = password;
        }
        public string ID { get { return id; } set { id = value; } }
        public void AddAnswers(string[] answers)
        {
            this.answers = new string[answers.Length];
            int i = 0;
            foreach (string s in answers)
            {
                if (s == null) this.answers[i] = "";
                else this.answers[i] = s;
                i++;
            }
        }
    }

    

    public class Question
    {
        //раздел. 1 - listening, 2 - reading, 3 - writing
        private int section;
        //текст вопроса
        private string text;


        public Question(string question, int section)
        {
            if (text == "") throw new ArgumentNullException();
            if ((section > 3) || (section < 0)) throw new ArgumentOutOfRangeException();
            
            this.text = question;
            this.section = section;
        }
        public string Text { get { return text; } }
        public int Section { get { return section; } }
    }
}
