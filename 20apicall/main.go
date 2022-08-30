package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("HTTP Server")
	PerformGetReq()  
	PerformPostJsonReq()
	PerformPostReqForm()
}


func PerformGetReq()  {
	const _url = "http://localhost:8000/get" 

	res,err := http.Get(_url) 

	if err != nil {
		panic(err)
	}

	defer res.Body.Close() 

	fmt.Println("status code : ",res.StatusCode) 
	fmt.Println("Content length : ",res.ContentLength) 
    
	content,_ := io.ReadAll(res.Body) 

	// fmt.Println(string(content))

	var resString strings.Builder 

	byteCount,_ := resString.Write(content)

	fmt.Println("Byte count is : ",byteCount) 
    fmt.Println("Response data : ",resString.String())
}


func PerformPostJsonReq(){
	var _url = "http://localhost:8000/post" 
	
	reqBody := strings.NewReader(`
	   {
		"name":"Mohammed Shahul",
		"age":27,
		"place":"malappuram"
	   }
	`)
   
    res,err :=	http.Post(_url,"application/json",reqBody)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	content,_ := io.ReadAll(res.Body) 

	fmt.Println(string(content))

}



func PerformPostReqForm()  {
	const _url = "http://localhost:8000/post" 

	data := url.Values{} 

	data.Add("name","Mohammed Shahul")
	data.Add("age","27") 
	data.Add("place","malappuram") 

 	res,err := http.PostForm(_url,data) 

	if err != nil {
		panic(err) 
	}

    defer res.Body.Close()
	
	content,_ := io.ReadAll(res.Body) 

	fmt.Println(string(content))
}