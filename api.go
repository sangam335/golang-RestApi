package main
import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"

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
func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w,"Homepage Endpoint hit")
}
func handleRequests(){
	http.HandleFunc("/",homePage)
	http.HandleFunc("/articles",allArticles)
	log.Fatal(http.ListenAndServe(":8080",nil))

}
func main(){
	handleRequests()
}
