package main

import (
"fmt"
)

func main() {
  fmt.Println(string(1010))
  // array := []string{"22","33","44"}
  barray := []byte("a slice is nice")

  // for index,element := range array {
  //   fmt.Println(index)
  //   fmt.Println(element)
  // }

    for index,element := range barray {
    fmt.Println(index)
    fmt.Println(element)
    fmt.Println(string(element))
    fmt.Println(barray)
  }
}


// func main() {

//   var array [4]int 
//   secondarray := [...]int{3, 1, 5}

//   for index,element := range array {
//     fmt.Println(index)
//     fmt.Println(element)
//   }

//   fmt.Println(secondarray)
// }