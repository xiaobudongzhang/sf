import edu.princeton.cs.algs4.StdIn;
import edu.princeton.cs.algs4.StdOut;


public class UF {
    private int[] id;
    private int count;

    //init
    public UF(int N) {
        count = N;
        id = new int[N];
        for (int i = 0; i < N; i++)
        {
            id[i] = i;
        }
    }

    public int count() {
        return count;
    }

    public int find(int p)
    {
        return id[p];
    }

    public boolean connected(int p, int q)
    {
        return find(p) == find(q);
    }

    public void union(int p, int q)
    {
        int pId = find(p);
        int qId = find(q);

        if (pId == qId)
        {
            return;
        }
        for (int i = 0; i< id.length; i++)
        {
            if (id[i] == pId)
            {
                id[i] = qId;
            }
        }
        count--;
        return;
    }

    public static void main(String[] args)
    {
       int N = StdIn.readInt();
       UF uf = new UF(N);
       while (!StdIn.isEmpty())
       {

           int p = StdIn.readInt();
           int q = StdIn.readInt();
           if (uf.connected(p, q))
           {
               continue;
           }
           uf.union(p, q);
           StdOut.println(p + " " + q);
           StdOut.println(uf.count() + " components");
       }
    }
}
