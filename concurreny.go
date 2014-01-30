package main

import (
  "fmt"
)

func main() {
    c := make(chan int)
    // Start three concurrent goroutines.  Numbers will be incremented
    go inc(0, c) 
    // go is a statement that starts a new goroutine.
    go inc(10, c)
    go inc(-805, c)
     fmt.Println(<-c, <-c, <-c)
}

func inc(i int, c chan int) {

    c <- i + 1 // <- is the "send" operator when a channel appears on the left.
    fmt.Println(i)
}