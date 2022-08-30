package main

import "fmt"

func main() {
	fmt.Print("---------- Functions ----------\n\n")
    hello()

	result :=addTwoNum(10,5) 
	fmt.Println("your result is ",result)

	result = addMoreNum(1,2,3,4,5,6,7,8,9,10)
	fmt.Println("your result is ",result)
}

func hello()  {
	fmt.Println("hello Go dev!..")
}


func addTwoNum(n1 int ,n2 int) int {
	return n1 + n2
}


func addMoreNum(values ...int)int{
	total :=0 
	for _,v := range values {
		total += v 
	}
	return total
}