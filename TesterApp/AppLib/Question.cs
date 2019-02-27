using System;

namespace AppLib
{

    public class Test
    {
        public Test(int id, BaseQuestions baseQuestions, Reading reading, Writing writing)
        {
            Id = id;
            BaseQuestions = baseQuestions;
            Reading = reading;
            Writing = writing;
        }

        public int Id { get; set; }
        public BaseQuestions BaseQuestions { get; set; }
        public Reading Reading { get; set; }
        public Writing Writing { get; set; }
    }


    public class Reading
    {
        public Reading(Question[] questions, string text)
        {
            Questions = new Question[questions.Length];
            questions.CopyTo(Questions, 0);
            Text = text;
        }
        public Question[] Questions { get; set; }
        public string Text { get; set; }
    }

    public class Writing
    {
        public Writing(string text)
        {
            Text = text;
        }

        public string Text { get; set; }
    }

    public class BaseQuestions
    {
        public BaseQuestions(Question[] questions)
        {
            Questions = new Question[questions.Length];
            questions.CopyTo(Questions, 0);
        }

        public Question[] Questions { get; set; }
    }

    public class Question
    {
        public Question(int id, string text, string optionA, string optionB,
            string optionC, string optionD)
        {
            if (id < 0) throw new ArgumentNullException();
            Id = id;
            Text = text;
            this.optionA = optionA;
            this.optionB = optionB;
            this.optionC = optionC;
            this.optionD = optionD;
        }

        public int Id { get; set; }
        public string Text { get; set; }
        public string optionA { get; set; }
        public string optionB { get; set; }
        public string optionC { get; set; }
        public string optionD { get; set; }

    }
}