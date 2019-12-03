using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.IO;
using data;
using System.Web.Script.Serialization;
using System.Net;
using System.Windows;
using System.Runtime.InteropServices;

namespace print
{
    class Program
    {
        static void Main(string[] args)
        {

            bool continuar = true;


            var lista = new List<Dados<dynamic>>();
            do
            {
                var adicio = new Dados<dynamic>();

                Console.Write("Digite o parâmetro: ");
                adicio.parametro = Console.ReadLine();
                Console.Write("Digite o valor: ");
                adicio.valor = Console.ReadLine();

                lista.Add(adicio);

                Console.Write("Deseja inserir mais campos? (S/N)");
                string continua = Console.ReadLine();

                if (continua == "S" || continua == "s")
                    continuar = true;
                else if (continua == "N" || continua == "n")
                    continuar = false;
                else
                    continuar = true;

            } while (continuar == true);



            var AccessToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiTmF0YWxpYSBQYXJzZXIiLCJpYXQiOjE1MTYyMzkwMjJ9.XHVYUrWeWX5FiAe2awgiMkJA0fwnDhCnug_3RHd3KVA";
            var request = WebRequest.Create("http://localhost:8080/api/usuarios/_create");
            var myHttpWebRequest = (HttpWebRequest)request;
            myHttpWebRequest.PreAuthenticate = true;
            myHttpWebRequest.Headers.Add("Authorization", "Bearer " + AccessToken);
            myHttpWebRequest.Accept = "application/json";
            request.ContentType = "application/json";
            request.Method = "POST";

            using (var streamWriter = new StreamWriter(request.GetRequestStream())){
                string json = new JavaScriptSerializer().Serialize(lista.Select(s => string.Format("{0} : {1}", s.parametro, s.valor)));
                streamWriter.WriteLine(json);
                streamWriter.Close();              
            }
        var response = (HttpWebResponse)request.GetResponse();
        using (var streamReader = new StreamReader(response.GetResponseStream()))
        {
                var result = streamReader.ReadToEnd();
        }     
    }

}
}

