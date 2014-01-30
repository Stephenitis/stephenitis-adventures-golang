package main

import (
	"net/http"
  "fmt"
  // "io/ioutil"
  "io"
  "os"
)

func main() {

resp, err := http.Get("http://google.com/")
if err != nil {
fmt.Println("ERROR")
}
defer resp.Body.Close()
fmt.Println(resp)
// body, err := ioutil.ReadAll(resp.Body)
out, err := os.Create("filename.html")
io.Copy(out, resp.Body)

}
