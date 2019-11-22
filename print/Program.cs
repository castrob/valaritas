using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.IO;
using data;
using System.Web.Script.Serialization;

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

            JavaScriptSerializer serializador = new JavaScriptSerializer();

            string stringObjSerializado = serializador.Serialize(lista.Select(s => string.Concat(s.parametro, ": ", s.valor)));


            Random random = new Random();
            string fileName = "json" + (random.Next(1000));
            string pathName = System.AppDomain.CurrentDomain.BaseDirectory.ToString();
            string path = pathName + "arq_json\\" + fileName + ".json";
            if (!File.Exists(path))
            {
                FileStream fs = File.Create(path);
                fs.Dispose();
            }

            StreamWriter stream = new StreamWriter(pathName + "arq_json\\"+ fileName + ".json");
            stream.WriteLine(stringObjSerializado);
            stream.Close();



        }
    }
}
