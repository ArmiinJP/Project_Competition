package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
   resp, err := http.Get("http://127.0.0.1:8080/avatar?email=abcd1234@gmail.com")
   if err != nil {
      log.Fatalln(err)
   }
   fmt.Println(resp.Header["Content-Type"])
   fmt.Println(resp.Status)
//We Read the response body on the line below.
   body, err := ioutil.ReadAll(resp.Body)
   if err != nil {
      log.Fatalln(err)
   }
//Convert the body to type string
   sb := string(body)
   log.Printf(sb)
}