using System;

namespace AppLib
{
    public class Question
    {
        public Question(string question, int section)
        {
            if (question == string.Empty) throw new ArgumentNullException();
            if (section > 3 || section < 0) throw new ArgumentOutOfRangeException();
            Text = question;
            Section = section;
        }

        //раздел. 1 - listening, 2 - reading, 3 - writing
        public string Text { get; }

        //текст вопроса
        public int Section { get; }
    }
}