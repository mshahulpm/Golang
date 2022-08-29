package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("welcome to study the time")
	currentTime := time.Now()
	fmt.Println(currentTime.Format("01-02-2006"))

	createdDate := time.Date(2021,time.April,10,12,30,0,0,time.UTC)
	fmt.Println(createdDate)
}