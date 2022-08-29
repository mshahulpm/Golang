package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Slices") 

	var fruits = []string {}
	fmt.Printf("type of fruits %T \n",fruits)

	fruits = append(fruits, "mango","apple","grapes") 
	fmt.Println(fruits)

	fruits = fruits[1:]              //append(fruits[1:])
	fmt.Println(fruits)              // [apple grapes]

	fruits = append(fruits[1:2])
	fmt.Println(fruits)              // [grapes]
	  
	highScores := make([]int,4) 

	highScores[0] = 234 
	highScores[1] = 334 
	highScores[2] = 434 
	highScores[3] = 834 
	// highScores[4] = 634    // this will give you an error because it is out of range  
    highScores = append(highScores, 232,322,355)  // this will not it will re-allocate memory  and size
    fmt.Println(highScores)

	sort.Ints(highScores) 
    fmt.Println(highScores)
    
	//  how to remove a value from slices based on index 
	var courses = []string {"js","ts","swift","python","css"} 
	fmt.Println(courses) 
	var index = 2 

	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses) 
}