using System.IO;
using System.Linq;
using System.Net;
using System.Text;
using Newtonsoft.Json;
using PainlessHttp.Serializer.JsonNet;
using RestSharp;


namespace AppLib
{
    public static class Server
    {
        public static Test GetTest(string id)
        {
            Test test;
            var client = new RestClient("http://138.68.78.205:8080/api/questions");
            var request = new RestRequest(Method.GET);
            request.AddHeader("Authorization", id);
            request.AddParameter("undefined", "undefined=", ParameterType.RequestBody);
            IRestResponse response = client.Execute(request);
            var serialized = Encoding.UTF8.GetString(response.RawBytes);
            using (var sw = new StreamWriter("test.json",
                false, Encoding.Default))
            {
                sw.Write(serialized);
                sw.Close();
            }

            // if response != OK => return null else return new object of class Test
            // for that decoding response.RawBytes (Bytes[]) to String
            // then forwarding that string in JsonConvert.DeserializeObject function
            return response.StatusCode != HttpStatusCode.OK ? null 
                : JsonConvert.DeserializeObject<Test>(Encoding.UTF8.GetString(response.RawBytes));
            
//            Test test;
//            var client = new RestClient("http://138.68.78.205:8080/api/questions");
//            var request = new RestRequest(Method.GET);
//            request.AddHeader("Authorization", id);
//            var answer = client.DownloadData(request);
//            using (FileStream fstream = new FileStream("test.json", FileMode.OpenOrCreate))
//            {
//                fstream.Write(answer, 0, answer.Length);
//                using (var sr = new StreamReader(fstream))
//                {
//                    string output = sr.ReadToEnd();
//                    try
//                    {
//                        test = JsonConvert.DeserializeObject<Test>(output);
//                    }
//                    catch(JsonException)
//                    {
//                        test = null;
//                    }
//                    sr.Close();
//                }
//            }
//            return test;
        }

        public static bool SendData(string id, TestAnswers testAnswers)
        {
            string answer = JsonConvert.SerializeObject(testAnswers);
            RestClient client = new RestClient("http://138.68.78.205:8080/api/answer");
            RestRequest request = new RestRequest(Method.POST);
            request.AddHeader("Content-Type", "application/json");
            request.AddHeader("Authorization", 
                id);
            request.AddParameter("undefined", answer, ParameterType.RequestBody);
            IRestResponse response = client.Execute(request);
            return response.StatusCode == HttpStatusCode.OK;
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