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
            q[0] = new Question("ASA", 0, 1, new string[0]);
            q[1] = new Question("ASP", 0, 1, new string[0]);
            q[2] = new Question("ASR", 2, 1, new string[2]{ "1", "2" });
            q[3] = new Question("AWA", 2, 2, new string[2] { "1", "2" });
            q[4] = new Question("Work it, make it, do it \nMakes us harder, better, faster, stronger!" +
                "\n N - n - now that don’t kill me \nCan only make me stronger...", 0, 3, new string[0]);
            return q;
        }
        public static bool SendData()
        {
            return true;
        }
    }
}
