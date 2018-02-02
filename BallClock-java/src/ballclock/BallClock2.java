package ballclock;

import java.util.Stack;
import java.util.Queue;
import java.util.IllegalFormatException;
import java.util.Iterator;
import java.util.LinkedList;

public class BallClock2 {
	Queue<Integer> min = new LinkedList<>();
	Queue<Integer> fiveMin = new LinkedList<>();
	Queue<Integer> hour = new LinkedList<>();
	Queue<Integer> reserve = new LinkedList<>();

	
	public static final int MINIMUM_RESERVE = 27;
	public static final int DEFAULT_RESERVE = 50;
	public static final int MAXIMUM_RESERVE = 127;
		
	public BallClock2(){
		new BallClock2(DEFAULT_RESERVE);
	}
	
	public BallClock2(int reserveSize){
		if(reserveSize < MINIMUM_RESERVE)
			throw new IllegalArgumentException("The ball reserve is too small.");
		if(reserveSize > MAXIMUM_RESERVE)
			throw new IllegalArgumentException("The ball reserve is too large.");
		
		for(int i = 0; i < reserveSize; i++){
			reserve.add(new Integer(i+1));
		}
	}
	
	public void moveBall() {
		min.add(reserve.remove());
		if(min.size() >= 5) {
			overflow(min, fiveMin);
			if(fiveMin.size() >= 12) {
				overflow(fiveMin, hour);
				if(hour.size() >= 12)
					overflow(hour, null);
			}
		}
	}
	
	public void overflow(Queue<Integer> overflowed, Queue<Integer> overflow) {
		if(overflow != null) {
			overflow.add(overflowed.remove());
		}
		
		while (!overflowed.isEmpty()) {
			reserve.add(overflowed.remove());
		}		
	}
	
	public String stackToString(Stack<Integer> stack) {
		StringBuilder sb = new StringBuilder();
		sb.append("[");
		
		if(stack.size() > 0) {
			sb.append(stack.get(0));
			
			for(int i = 1; i < stack.size(); i++) {
				sb.append(",");
				sb.append(stack.get(i));
			}
		}
				
		sb.append("]");
		
		return sb.toString();
	}
	
	public String queueToString(Queue<Integer> queue) {
		StringBuilder sb = new StringBuilder();
		sb.append("[");
		
		Iterator<Integer> i = queue.iterator();
		if(i.hasNext()) {
			sb.append(i.next());
		
			while(i.hasNext()) {
				sb.append(",");
				sb.append(i.next());
			}
		}
				
		sb.append("]");
		
		return sb.toString();
	}
	
	
	@Override
	public int hashCode() {
		final int prime = 31;
		int result = 1;
		result = prime * result
				+ ((fiveMin == null) ? 0 : fiveMin.hashCode());
		result = prime * result
				+ ((hour == null) ? 0 : hour.hashCode());
		result = prime * result
				+ ((min == null) ? 0 : min.hashCode());
		result = prime * result
				+ ((reserve == null) ? 0 : reserve.hashCode());
		return result;
	}
	
	@Override
	public boolean equals(Object obj) {
		if (this == obj)
			return true;
		if (obj == null)
			return false;
		if (getClass() != obj.getClass())
			return false;
		BallClock2 other = (BallClock2) obj;
				
		if(!min.equals(other.min))
			return false;
		if(!fiveMin.equals(other.fiveMin))
			return false;
		if(!hour.equals(other.hour))
			return false;
		if(!reserve.equals(other.reserve))
			return false;
		
		return true;
	}

	@Override
	public String toString() {
		return "{Min:" + queueToString(min) + ", FiveMin:" + queueToString(fiveMin) + ", Hour:" + queueToString(hour) + ", Main:" + queueToString(reserve) + "}";
	}
	
	public static void main(String[] args) {
		if(args.length == 1) {
			//Mode 1, determine how days pass before the clock returns to its initial ordering.
			try {
				int numberOfBalls = Integer.parseInt(args[0]);
				if(numberOfBalls < 27) {
					System.out.println("The number of balls but be 27 or greater.");	
				} else if(numberOfBalls > 127) {
					System.out.println("The number of balls but be 127 or less.");	
				} else {
					//Start time tracking
					long start = System.currentTimeMillis();
					BallClock2 ballClock = new BallClock2(numberOfBalls);
					BallClock2 initialClock = new BallClock2(numberOfBalls);
					int minutesElapsed = 0;
					do {
						minutesElapsed++;
						ballClock.moveBall();
					} while(!ballClock.equals(initialClock));
					double days = Math.ceil(minutesElapsed / 60.0 / 24.0 );
					long end = System.currentTimeMillis();
					
					System.out.println(numberOfBalls + " balls cycle after " + days + " days.");
					System.out.println("Completed in " + (end - start) / 1000.0 + " seconds");
				}
			} catch (IllegalFormatException e) {
				System.out.println("The first parameters but be the number of balls.");
			}
		} else if(args.length == 2) {
			//Mode 2, determine position of the balls after the specified number of minutes.
			try {
				int numberOfBalls = Integer.parseInt(args[0]);
				int numberOfminutes = Integer.parseInt(args[1]);
				if(numberOfBalls < 27) {
					System.out.println("The number of balls but be 27 or greater.");	
				} else if(numberOfBalls > 127) {
					System.out.println("The number of balls but be 127 or less.");	
				} else {
					//Start time tracking
					BallClock2 ballClock = new BallClock2(numberOfBalls);
					int minutesElapsed = 0;
					do {
						minutesElapsed++;
						ballClock.moveBall();
					} while(minutesElapsed < numberOfminutes);
					
					System.out.println(ballClock.toString());
				}
			} catch (NumberFormatException e) {
				System.out.println("The first parameters but be the number of balls.");
			}
		} else {
			System.out.println("This program must have exactly one or two arguments.");
		}
		
	}
}

