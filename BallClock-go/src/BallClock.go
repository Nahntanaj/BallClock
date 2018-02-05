package main

import (
  "fmt"
  "bytes"
  "os"
  "strconv"
)

type BallClock struct { 
	min []int
	fiveMin []int 
	hour []int
	reserve []int
}

func overflow(overflowed []int, overflow []int, reserve []int) ([]int, []int, []int){
	overflow = append(overflow, overflowed[len(overflowed) - 1])
	overflowed, reserve = spill(overflowed[0:len(overflowed) - 2], reserve)
  
  return overflowed, overflow, reserve
}

func spill(overflowed []int, reserve []int) ([]int, []int){
	for i:=len(overflowed) - 1;i>=0;i-- {
		reserve = append(reserve, overflowed[i])
	}
  overflowed = make([]int,0)
  return overflowed, reserve
}

func incrementMinute(clock BallClock) BallClock{
  fmt.Println(toString(clock))
	clock.min = append(clock.min, clock.reserve[0])
	clock.reserve = clock.reserve[1:]
	if len(clock.min) > 4 {
		clock.min, clock.fiveMin, clock.reserve = overflow(clock.min, clock.fiveMin, clock.reserve)
		if len(clock.fiveMin) > 11 {
			clock.fiveMin, clock.hour, clock.reserve = overflow(clock.fiveMin, clock.hour, clock.reserve)
			if len(clock.hour) > 11 {
				clock.hour, clock.reserve = spill(clock.hour, clock.reserve)
			}
		}	
	}
  return clock
}

func toString(clock BallClock) string {
  var buffer bytes.Buffer
  buffer.WriteString("{")
  
  //Mins
  buffer.WriteString("\"Min\":")
  buffer.WriteString("[")
  if len(clock.min) > 0 {
    buffer.WriteString(strconv.Itoa(int(clock.min[0])))
    for i:=1; i<len(clock.min); i++ {
      buffer.WriteString(",")
      buffer.WriteString(strconv.Itoa(int(clock.min[i])))
    }
  }
  buffer.WriteString("]")
  
  //FiveMins
  buffer.WriteString(",\"FiveMin\":")
  buffer.WriteString("[")
  if len(clock.fiveMin) > 0 {
    buffer.WriteString(strconv.Itoa(int(clock.fiveMin[0])))
    for i:=1; i<len(clock.fiveMin); i++ {
      buffer.WriteString(",")
      buffer.WriteString(strconv.Itoa(int(clock.fiveMin[i])))
    }
  }
  buffer.WriteString("]")
  
  //Hours
  buffer.WriteString(",\"Hour\":")
  buffer.WriteString("[")
  if len(clock.hour) > 0 {
    buffer.WriteString(strconv.Itoa(int(clock.hour[0])))
    for i:=1; i<len(clock.hour); i++ {
      buffer.WriteString(",")
      buffer.WriteString(strconv.Itoa(int(clock.hour[i])))
    }
  }
  buffer.WriteString("]")
  
  //Main
  buffer.WriteString(",\"Main\":")
  buffer.WriteString("[")
  if len(clock.reserve) > 0 {
    buffer.WriteString(strconv.Itoa(int(clock.reserve[0])))
    for i:=1; i<len(clock.reserve); i++ {
      buffer.WriteString(",")
      buffer.WriteString(strconv.Itoa(int(clock.reserve[i])))
    }
  }
  buffer.WriteString("]")
  
  buffer.WriteString("}")
  return buffer.String()
}

func equals(clock1, clock2 BallClock) bool{
  if len(clock1.min) != len(clock2.min) {
    return false
  }
  
  if len(clock1.fiveMin) != len(clock2.fiveMin) {
    return false
  }
  
  if len(clock1.hour) != len(clock2.hour) {
    return false
  }
  
  if len(clock1.reserve) != len(clock2.reserve) {
    return false
  }
  
  for i:=0;i<len(clock1.min);i++ {
    if(clock1.min[i] != clock2.min[i]){
      return false
    }
  }
  
  for i:=0;i<len(clock1.fiveMin);i++ {
    if(clock1.fiveMin[i] != clock2.fiveMin[i]){
      return false
    }
  }
  
  for i:=0;i<len(clock1.hour);i++ {
    if(clock1.hour[i] != clock2.hour[i]){
      return false
    }
  }
  
  for i:=0;i<len(clock1.reserve);i++ {
    if(clock1.reserve[i] != clock2.reserve[i]){
      return false
    }
  }
  
  return true
}


func main() {
  args := os.Args
  if(len(args) == 2) {
    reserveCapacity, err := strconv.Atoi(args[1])
    
    if err != nil {
      os.Exit(-1)
    }
    
    if(reserveCapacity < 27){
      os.Exit(-1)
    } else if(reserveCapacity > 127){
      os.Exit(-1)
    } else {
      var reserve []int
      for i := 1; i< reserveCapacity; i++ {
        reserve = append(reserve, i)
      }
      ballclock := BallClock{[]int{}, []int{}, []int{}, reserve}
      ballclockOriginal := BallClock{[]int{}, []int{}, []int{}, reserve}
      minutesRan := 1
      ballclock = incrementMinute(ballclock)
      for ;!equals(ballclock, ballclockOriginal); {
        minutesRan++
        ballclock = incrementMinute(ballclock)
      }
      
      fmt.Print( "%d balls cycled after %d days.", reserveCapacity, minutesRan / 720)
    }
  } else if (len(args) == 3 ) {
    reserveCapacity, err := strconv.Atoi(args[1])
    minutesToRun, err := strconv.Atoi(args[2])
    
    if err != nil {
      os.Exit(-1)
    }
    
    if(reserveCapacity < 27){
      os.Exit(-1)
    } else if(reserveCapacity > 127){
      os.Exit(-1)
    } else {
      var reserve []int
      for i := 1; i< reserveCapacity; i++ {
        reserve = append(reserve, (int)(i))
      }
      ballclock := BallClock{[]int{}, []int{}, []int{}, reserve}
      for i := 0; i< minutesToRun; i++ {
        ballclock = incrementMinute(ballclock)
      }
      
      fmt.Println(toString(ballclock))
    }
  }
}

