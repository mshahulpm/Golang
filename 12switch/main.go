package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch")

	rand.Seed(time.Now().Unix())
	diceNo := rand.Intn(6)+1 
	// fmt.Println("Value of dice is ",diceNo)

	switch diceNo {
	case 1 :
		fmt.Println("hey 1")
	case 2 :
		fmt.Println("hey 2")
	case 3 :
		fmt.Println("hey 3")
	case 4 :
		fmt.Println("hey 4")
	case 5 :
		fmt.Println("hey 5")
	case 6 :
		fmt.Println("hey 6")
	default:
		fmt.Println("default")
	}
}