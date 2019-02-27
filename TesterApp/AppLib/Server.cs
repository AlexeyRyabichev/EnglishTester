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
            Test test;
            using (var sr = new StreamReader(new FileStream("test.json", FileMode.Open)))
            {
                string output = sr.ReadToEnd();
                test = JsonConvert.DeserializeObject<Test>(output);
                sr.Close();
            }
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