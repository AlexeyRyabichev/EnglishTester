using System.IO;
using System.Linq;
using System.Net;
using System.Text;
using Newtonsoft.Json;
using RestSharp;

namespace AppLib
{
    public static class Server
    {
        public static Test GetTest()
        {
            Question question1 = new Question(1, "asdkjhgfdshjkl", "a", "b", "c", "d");
            Question question2 = new Question(2, "AAAAAAAAAAAAAAA", "a", "b", "c", "d");
            Question[] questions = new Question[2] { question1, question2 };
            Reading reading = new Reading(questions, "ITS READING TIME");
            Writing writing = new Writing("xdyuyihojp\nzrxtcfyvgubh\nzrxtcyvgbh\nerxtcyvubhnj\nhj");
            BaseQuestions baseQuestions = new BaseQuestions(questions);
            Test test = new Test(1, baseQuestions, reading, writing);
            using (var sw = new StreamWriter("test1.json",
                false, Encoding.Default))
            {
                var serialized = JsonConvert.SerializeObject(test);
                sw.Write(serialized);
                sw.Close();
            }
            /*using (var sr = new StreamReader(new FileStream("test.json", FileMode.Open)))
            {
                string output = sr.ReadToEnd();
                test = JsonConvert.DeserializeObject<Test>(output);
                sr.Close();
            }*/
            return test;
        }

        public static bool SendData(Student student)
        {
            using (var sw = new StreamWriter("student.json",
                false, Encoding.Default))
            {
                var serialized = JsonConvert.SerializeObject(student);
                sw.Write(serialized);
                sw.Close();
            }

            return true;
        }

        public static string Authentication(string email, string password)
        {
            var client = new RestClient("http://138.68.78.205:8080/api/login");
            var request = new RestRequest(Method.POST);
            email = email.Replace("@", "%40");
            request.AddHeader("Content-Type", "application/x-www-form-urlencoded");
            request.AddParameter("undefined", "email=" + email + "&password=" + password,
                ParameterType.RequestBody);
            var response = client.Execute(request);
            return response.StatusCode == HttpStatusCode.OK
                ? response.Headers.ToList().Find(x => x.Name == "Authorization").Value.ToString()
                : string.Empty;
        }
    }
}