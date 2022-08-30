package main

import "fmt"

func main() {
	fmt.Println("------- Methods ------")  

	user := User{"shahul","shahul@tec.dev",true,27} 
	fmt.Println(user)
	fmt.Printf("user details are %+v\n",user)
	user.GetStatus()
	user.newEmail("new@tec.dev")
	fmt.Println(user.Email)
}


type User struct {
	Name    string 
	Email   string 
	Status  bool 
	Age     int 
}

func (u User) GetStatus() {
	fmt.Println("is user is active : ",u.Status)
}

func (u User) newEmail(mail string)  {
	fmt.Println("Old email of the user : ",u.Email)
	u.Email =  mail  //"newmail@tec.dev"
	fmt.Println("New email of the user : ",u.Email)
}