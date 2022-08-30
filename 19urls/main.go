package main

import (
	"fmt"
	"net/url"
)

const myUrl  = "https://hello.dev:3000/learn?course=reactjs&payment=300" 

func main() {
	fmt.Println("------- URLS -------")

	result,_ := url.Parse(myUrl) 

	fmt.Println(result.Scheme) 
	fmt.Println(result.Query()) 
	fmt.Println(result.Host) 
	fmt.Println(result.Path) 
	fmt.Println(result.RawQuery) 
	fmt.Println(result.Port()) 

	quries := result.Query() 

	for key,val := range quries {
		fmt.Printf("query is %v : %v\n",key,val) 
	}


	partsOfUrl := &url.URL{
		Scheme: "https",
		Host: "hello.dev",
		Path: "/learn",
		RawPath: "user=shahul",
	}

	_url := partsOfUrl.String() 

	fmt.Println(_url)
}