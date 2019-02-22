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
            Question[] q = new Question[3];
            q[0] = new Question("ASA", 0, 1, new string[0]);
            q[1] = new Question("ASP", 0, 1, new string[0]);
            q[2] = new Question("ASR", 2, 1, new string[0]);
            return q;
        }
        public static bool SendData()
        {
            return true;
        }
    }
}
