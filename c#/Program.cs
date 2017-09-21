using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.Threading;

namespace rudp {
    class Program {
        static int index = 0;

        static string dumpByte(byte[] data, int len) {
            string s = "";
            if (len == 0) {
                return s;
            }
            for (int i = 0; i < len - 1; ++i) {
                s += data[i] + ", ";
            }
            s += data[len - 1];

            return s;
        }

        static void dump(List<PackageBuffer> pkg_queue) {
            for (int i = 0; i < pkg_queue.Count; ++i) {
                Console.Out.Write(string.Format("pkg[ {0} ]: {1}\n", index, dumpByte(pkg_queue[i].buffer, pkg_queue[i].len)));
                index++;
            }
        }

        static void Main(string[] args) {
            Rudp U = new Rudp();

            byte[] d1 = { 1, 2, 3, 4 };
            byte[] d2 = { 5, 6, 7, 8 };
            byte[] d3 = { 3, 0, 0, 0 };
            byte[] d4 = { 3, 0, 2, 2 };

            U.Send(d1, 4);
            U.Send(d2, 3);
            dump(U.Update());

            dump(U.Update());

            U.Update(d3, 4);
            dump(U.Update(d4, 4));

            while (true) {
                Thread.Sleep(1000);
            }
        }
    }
}
