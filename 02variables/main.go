package main

import "fmt"

const LoginToken string = "blahblahtoken" // variable with capital letter is public

func main() {
	var username string = "Mohammed Shahul"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n",username) 

	var isLoggedIn bool = false 
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n",isLoggedIn) 

	var smallvalue uint8 = 255
	fmt.Println(smallvalue)
	fmt.Printf("Variable is of type: %T \n",smallvalue) 

	var smallFloat float64 = 255.45454545454545555555555512541254125425412
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n",smallFloat) 

	// default values and some aliases 
	var anotherVariable int 
	fmt.Println(anotherVariable) 
	fmt.Printf("Variable is of type: %T \n",anotherVariable) 

	// no var style 
	no_of_people := 129 
	fmt.Println(no_of_people)
	fmt.Printf("Variable is of type: %T \n",no_of_people) 
	
}