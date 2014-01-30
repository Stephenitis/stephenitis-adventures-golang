package main

import (
  "fmt"
)

var bottlesOfBeer int = 99

func main() {
  say(bottlesOfBeer)
}

func say(num int) () {
  removeBeer()
  fmt.Println(fmt.Sprintf("%d bottles of beer", num))
}

func removeBeer() {
  bottlesOfBeer--
  if bottlesOfBeer >= 1 {
    say(bottlesOfBeer)
  } else if bottlesOfBeer < 1 {
    fmt.Println("done")
  }
}