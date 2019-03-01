using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace AppLib
{
    public class TestAnswers
    {

        public TestAnswers() { }

        public VariableAnswer[] Base { get; set; }
        public VariableAnswer[] Reading { get; set; }
        public string Writing { get; set; }
    }

    public class VariableAnswer
    {
        public VariableAnswer(int id, string answer)
        {
            Id = id;
            Answer = answer;
        }
        public int Id { get; set; }
        public string Answer { get; set; }
    }
}
