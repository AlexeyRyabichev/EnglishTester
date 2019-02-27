namespace AppLib
{
    public class Student
    {
        public Student(string email, string password, string id)
        {
            Email = email;
            Password = password;
            ID = id;
        }

        public string Email { get; }
        public string Password { get; }
        public string[] Answers { get; private set; }
        public string ID { get; set; }

        public void AddAnswers(string[] answers)
        {
            Answers = new string[answers.Length];
            var i = 0;
            foreach (var s in answers)
            {
                if (s == null) Answers[i] = string.Empty;
                else Answers[i] = s;
                i++;
            }
        }
    }
}