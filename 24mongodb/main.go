package main

import (
	"fmt"
	"log"
	"mongodbapi/routers"
	"net/http"
)

func main() {
	fmt.Println("Mongodb database ")
	fmt.Println("Server is running on port 4000 ")

	r := routers.Router() 

	log.Fatal(http.ListenAndServe(":4000",r)) 
	
}