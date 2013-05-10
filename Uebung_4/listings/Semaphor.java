import java.util.concurrent.BlockingQueue;
import java.util.concurrent.LinkedBlockingQueue;

public class Semaphor {
	int count;
	BlockingQueue<Thread> queue;

	public Semaphor(int initialcount){
		count = initialcount;
		queue = new LinkedBlockingQueue<Thread>();
	}

	public void P(){
		synchronized(this){
			if(count==0){
				try{
					queue.put(Thread.currentThread());
					wait();
				}catch(InterruptedException ie){
					//this should never happen
					System.err.println("caught InterruptedException in wait()");
				}
			}
			count--;
		}
	}

	public void V(){
		synchronized(this){
			count++;
			if(count==1) {
				queue.poll().notify();
			}
		}
	}
}
