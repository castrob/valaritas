using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace data
{
    public class Dados<TValue>
    {
        public TValue parametro { get; set; }
        public TValue valor { get; set; }

        public Dados()
        {

        }
    }
}

