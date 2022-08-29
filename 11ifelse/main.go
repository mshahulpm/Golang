package main

import "fmt"

func main() {
	fmt.Println("If Else")
     
	count := 11 
	var result string

	if count <10 {
        result = "Regular "
	}else if count > 10 {
		result = "he he user"
	}else {
		result ="exact 10"
	}
	fmt.Println(result)

	if num :=3 ; num<10 {
		fmt.Println("Num is less than 10")
	}else {
		fmt.Println("Num is not less than 10")
	}
}