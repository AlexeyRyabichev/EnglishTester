using System;

namespace AppLib
{

    public class Test
    {
        public Test(int id, Questions[] baseQuestions, Reading reading, string writing)
        {
            Id = id;
            BaseQuestions = baseQuestions;
            Reading = reading;
            Writing = writing;
        }

        public int Id { get; set; }
        public Questions[] BaseQuestions { get; set; }
        public Reading Reading { get; set; }
        public string Writing { get; set; }
    }


    public class Reading
    {
        public Reading(Questions[] questions, string text)
        {
            Questions = new Questions[questions.Length];
            questions.CopyTo(Questions, 0);
            Question = text;
        }
        public Questions[] Questions { get; set; }
        public string Question { get; set; }
    }


    public class Questions
    {
        public Questions(int id, string text, string optionA, string optionB,
            string optionC, string optionD)
        {
            if (id < 0) throw new ArgumentNullException();
            Id = id;
            Question = text;
            this.optionA = optionA;
            this.optionB = optionB;
            this.optionC = optionC;
            this.optionD = optionD;
        }

        public int Id { get; set; }
        public string Question { get; set; }
        public string optionA { get; set; }
        public string optionB { get; set; }
        public string optionC { get; set; }
        public string optionD { get; set; }

    }
}