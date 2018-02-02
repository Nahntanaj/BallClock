package src

import (

)

type BallClock struct { 
	min []uint8
	fiveMin []uint8 
	hour []uint8
	reserve = []uint8
}

func overflow(overflowed []uint8, overflow []uint8){
	overflow = append(overflow, overflowed[:len(overflowed) - 1])
	spill(overflowed[0:len(overflowed) - 2]
}

func spill(overflowed []uint8){
	for i=len(overflowed) - 1;i>=0;i-- {
		reserve = append(overflowed[i]);
		overflowed = []uint8;
	}
}

func incrementMinute(){
	min = append(min, reserve[0])
	reserve = reserve[1:]
	if len(min) > 4 {
		overflow(min, fiveMin)
		if len(fiveMin) > 11 {
			overflow(fiveMin, hour)
			if len(hour) > 11 {
				spill(hour)
			}
		}	
	}
}

func main(reserveCapacity uint8) string{
	if(reserveCapacity < 27){
		//error message
	} else if(reserveCapacity > 127){
		//error message
	} else {
		ballclock = BallClock([5]uint8, [12]uint8, [12]uint8, [reserveCapacity]uint8)
		 	
	}
		
		
	
		
}

