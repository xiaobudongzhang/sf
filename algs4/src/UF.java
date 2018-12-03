public class UF
{
    private int[] id;
    private int count;
    //init
    public UF(int N)
    {
        count = N;
        id = new int[N];
        for (int i = 0; i < N; i++)
        {
            id[i] = i;
        }
    }

    public static void main(String[] args)
    {
       System.out.println("xxxxxxxxxxxxx");
    }
}
