package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("----Channels and Deadlocks")

	myChan := make(chan int)
	wg := &sync.WaitGroup{}
	// myChan <- 5

	// fmt.Println(<-myChan)
	wg.Add(2)
	// go func(ch chan int, wg *sync.WaitGroup) {
	// 	fmt.Println(<-ch)
	// 	fmt.Println(<-ch)
	// 	wg.Done()
	// }(myChan, wg)

	// go func(ch chan int, wg *sync.WaitGroup) {
	// 	ch <- 5
	// 	ch <- 6
	// 	close(ch)
	// 	wg.Done()
	// }(myChan, wg)

	go func(ch <-chan int, wg *sync.WaitGroup) {
		// close(ch)  can`t do that because receive only
		val, isChannelOpen := <-ch
		fmt.Println(isChannelOpen, val)
		wg.Done()
	}(myChan, wg)

	go func(ch chan int, wg *sync.WaitGroup) {
		close(ch)
		wg.Done()
	}(myChan, wg)

	wg.Wait()

}
