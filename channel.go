package main

import (
  "fmt"
  "time"
)

func pinger(c chan string) {
  for i := 0; ; i++ {
    // c <- "ping" means send "ping"
    c <- "ping"
  }
}

func ponger(c chan string) {
  for i := 0; ; i++ {
    c <- "pong"
  }
}

func printer(c chan string) {
  for {
    /*
     means receive a message and store it in msg.
     When pinger attempts to send a message on the channel 
     it will wait until printer is ready to receive the 
     message. (this is known as blocking)
    */
    msg := <- c
    fmt.Println(msg)
    time.Sleep(time.Second * 1)
  }
}

func main() {
  var c chan string = make(chan string)

  go printer(c)
  go pinger(c)
  go ponger(c)
  /*
   With a goroutine we return immediately to the next line 
   and don't wait for the function to complete. This is why 
   the call to the Scanln function has been included; without it 
   the program would exit before being given the opportunity to print all the numbers.
  */
  var input string
  fmt.Scanln(&input)
}