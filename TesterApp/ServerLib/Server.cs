using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using TesterLib;
using Newtonsoft.Json;
using System.IO;
using RestSharp;


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
            using (StreamWriter sw = new StreamWriter("student.json", 
                false, System.Text.Encoding.Default))
            {
                
                string serialized = JsonConvert.SerializeObject(student);
                sw.Write(serialized);
                sw.Close();
            }
            return true;
        }
        public static string Authentication(string email, string password)
        {
            var client = new RestClient("http://138.68.78.205:8080/api/login");
            var request = new RestRequest(Method.POST);
            email.Replace("@", "%40");
            request.AddHeader("Content-Type", "application/x-www-form-urlencoded");
            request.AddParameter("undefined", "email="+ email +"&password="+password, 
                ParameterType.RequestBody);
            IRestResponse response = client.Execute(request);
            if (response.StatusCode == System.Net.HttpStatusCode.Unauthorized)
                return "";
            else if (response.StatusCode == System.Net.HttpStatusCode.OK)
                return response.Headers.ToList().Find(x => x.Name == "Authorization").Value.ToString();

            return "";
        }
    }
}
