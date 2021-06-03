package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)
type Article struct{
	Title string `json:"Title"`
	Desc string `json:"Title"`
	Content string `json:"Title"`
}
type Articles []Article
func allArticles(w http.ResponseWriter,r *http.Request){
	articles:=Articles{
		Article{Title:"Test Title",Desc:"Test description",Content:"hello world"},

	}
	fmt.Println("Endpoint Hit:All Article Endpoint")
	json.NewEncoder(w).Encode(articles)

}
func testPostArticles(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w,"test POST endpoint worked")
}
func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w,"Homepage Endpoint hit")
}
func handleRequests(){
	myRouter:=mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/",homePage)
	myRouter.HandleFunc("/articles",allArticles).Methods("GET")
	myRouter.HandleFunc("/articles",testPostArticles).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8080",myRouter))

}
func main(){
	handleRequests()
}
