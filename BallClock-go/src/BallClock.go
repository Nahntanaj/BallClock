package src

import (
  "fmt"
  "buffer"
  "os"
  "strconv"
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

func toString() string {
  var buffer bytes.Buffer
  buffer.writeString("{")
  
  //Mins
  buffer.writeString("\"Min\":")
  buffer.writeString("[")
  if len(min) > 0 {
    buffer.writeString(min[0])
    for i:=1; i<len(min); i++ {
      buffer.writeString(",")
      buffer.writeString(min[i])  
    }
  }
  buffer.writeString("]")
  
  //FiveMins
  buffer.writeString(",\"FiveMin\":")
  buffer.writeString("[")
  if len(fiveMin) > 0 {
    buffer.writeString(fiveMin[0])
    for i:=1; i<len(fiveMin); i++ {
      buffer.writeString(",")
      buffer.writeString(fiveMin[i])  
    }
  }
  buffer.writeString("]")
  
  //Hours
  buffer.writeString(",\"Hour\":")
  buffer.writeString("[")
  if len(hour) > 0 {
    buffer.writeString(hour[0])
    for i:=1; i<len(hour); i++ {
      buffer.writeString(",")
      buffer.writeString(hour[i])  
    }
  }
  buffer.writeString("]")
  
  buffer.writeString("}")
  return buffer.String()
}

func main() {
  args := os.Args
  
  if(len(args) == 2) {
    reserveCapacity, err := Atoi(args[1])
    
    if(reserveCapacity < 27){
      os.Exit(-1)
    } else if(reserveCapacity > 127){
      os.Exit(-1)
    } else {
      ballclock = BallClock([5]uint8, [12]uint8, [12]uint8, [reserveCapacity]uint8)
    }
  } else if (len(args) == 3 ) {
    reserveCapacity, err := Atoi(args[1])
    minutesToRun, err := Atoi(args[2])
    
    if err != nil {
      os.Exit(-1)
    }
    
    if(reserveCapacity < 27){
      os.Exit(-1)
    } else if(reserveCapacity > 127){
      os.Exit(-1)
    } else {
      ballclock = BallClock([5]uint8, [12]uint8, [12]uint8, [reserveCapacity]uint8)
    }
  }
		
		
	
		
}

