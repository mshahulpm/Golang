package main

import "fmt"

func main() {
	fmt.Println("Struct")  

	user := User{"shahul","shahul@tec.dev",true,27} 
	fmt.Println(user)
	fmt.Printf("user details are %+v\n",user)
}


type User struct {
	Name    string 
	Email   string 
	Status  bool 
	Age     int 

}