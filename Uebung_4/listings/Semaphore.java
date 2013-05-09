public class Semaphore {
	private int zaehler = 0;
    private Thread warteschlange = null;
    public void init(int i)
    {
        zaehler = i;
        warteschlange = null;
    }
    public void P(Thread thread)
    {
        zaehler--;
        if(zaehler<0)
        {
            warteschlange = thread;
            try
            {
              warteschlange.suspend();
            }
            catch(Exception e)
            {
            }
        }
    }
    public void V()
    {
        zaehler++;
//	      if(warteschlange!=null)
        if(zaehler<=0)
            warteschlange.resume(); 
    }
}

