package main

import (
	"fmt"
	"net/http"
	"sync"
)

// func main() {
// 	fmt.Println("---------Go routine--------")
// 	go greeter("Hello")
// 	greeter("World")
// }

// func greeter(s string){
// 	for i :=0;i<6;i++{
// 		time.Sleep(2*time.Millisecond)
// 	fmt.Println(s)
// 	}
// }


var wg sync.WaitGroup 
var mut sync.Mutex
var signals = []string{"test"} 

func  main()  {

	websiteList := []string{
		"https://lco.dev",
		"https://go.dev",
		"https://google.com",
		"https://fb.com",
		"https://github.com",
	}

	for _,web := range websiteList {
      go  getStatusCode(web) 
	  wg.Add(1)
	}

	wg.Wait()
	fmt.Println(signals)
}


func getStatusCode(endpoint string)  {

	defer wg.Done()

	res,err := http.Get(endpoint) 

	if err != nil {
		fmt.Println("Error in end point")
	}else {
		mut.Lock()
		signals = append(signals, endpoint) 
		mut.Unlock()
	}

	fmt.Printf("%d status code for %s \n",res.StatusCode,endpoint)  
}