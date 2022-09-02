package main

import (
	"fmt"
	"math/big"

	// "math/rand"
	"crypto/rand"
)

func main() {
	fmt.Println("------Math and Random------")

	// random no using math.rand
	// rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(10))

	// random no using crypto.rand

	my_rand, _ := rand.Int(rand.Reader, big.NewInt(5))

	fmt.Println(my_rand)

}
