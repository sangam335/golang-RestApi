package main

import (
	 "fmt"
	 "log"
	 "net/http"
)

func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "Hello world")
}
func handleRequests(){
	http.HandleFunc("/",homePage)
  log.Fatal(http.ListenAndServe(":8080",nil))
}
func main(){
	handleRequests()
}
