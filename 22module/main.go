package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)


func main() {
	fmt.Println("------- Module -------")
	greetings()

	r := mux.NewRouter()
	r.HandleFunc("/",serveHome).Methods("GET")

	log.Fatal(http.ListenAndServe(":4000",r)) 
}


func greetings()  {
	fmt.Println("Hello there...") 
}


func serveHome(w http.ResponseWriter,r *http.Request)  {
	 w.Write([]byte("<h1>Hello welcome to app</h1>"))
}