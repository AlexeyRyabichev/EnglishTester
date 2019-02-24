using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using TesterLib;

namespace ServerLib
{
    public static class Server
    {
        public static Question[] GetQuestions()
        {
            Question[] q = new Question[5];
            q[0] = new Question("ASA", 1);
            q[1] = new Question("ASP", 1);
            q[2] = new Question("ASR", 1);
            q[3] = new Question("AWA", 2);
            q[4] = new Question("Work it, make it, do it \nMakes us harder, better, faster, stronger!" +
                "\n N - n - now that don’t kill me \nCan only make me stronger...", 3);
            return q;
        }
        public static bool SendData(Student student)
        {
            return true;
        }
        public static string Authentication(string email, string password)
        {
            return "0";
        }
    }
}
