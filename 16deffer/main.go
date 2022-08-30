package main

import "fmt"

func main() {
	defer fmt.Println("hello ")           // 4
	defer fmt.Println("world ")           // 3
	defer fmt.Println("world2 ")          // 2
	fmt.Println("------ deffer -------")  // 1
	myDeffer()
}




func myDeffer()  {
	for i :=0 ;i<5;i++ {
		defer fmt.Println(i)
	}
}