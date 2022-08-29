package main

import "fmt"

func main() {
	fmt.Println("Arrays") 

	var names [5]string 

	names[0] = "shahul"
	names[1] = "myran"
	names[2] = "shameem"
	names[4] = "nazeem"

	fmt.Println("Names list ",names,len(names))

	var letters = [3]string {"a","b","c"} 
	fmt.Println("List of letters ",letters)
}