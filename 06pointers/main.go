package main

import "fmt"

func main() {
	fmt.Println("Pointers")

	// var ptr *int 
	// fmt.Println("Value of pointer is ",ptr)

	myNumber := 23 

	var ptr2 = &myNumber 

	fmt.Println("Value of pointer is ",ptr2) // memory address 
	fmt.Println("Value of pointer is ",*ptr2) // actual value 

	*ptr2 = *ptr2 * 2   // this will manipulate actual value 
	fmt.Println(myNumber,*ptr2)
}