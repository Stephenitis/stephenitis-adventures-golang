package main

import (
	"net/http"
  "fmt"
  "io/ioutil"
)

func main() {

resp, err := http.Get("http://google.com/")
if err != nil {
fmt.Println("ERROR")
}
defer resp.Body.Close()
body, err := ioutil.ReadAll(resp.Body)

fmt.Println(string(body))

}
