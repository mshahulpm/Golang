package main

import "fmt"

func main() {
	fmt.Println("        Loops         ")
	fmt.Println("______________________")

	letters := []string{"a","b","c","d","e","f","g","h","i"}

	fmt.Println(letters)

	// for i:=0;i<len(letters);i++{
	// 	fmt.Println(letters[i])
	// }

	// for i:= range(letters) {
	// 	fmt.Println(letters[i]) 
	// }

	// for index,lt := range letters {
	// 	fmt.Println(index," : ",lt)
	// }

	i := 1 
	for i<10 {

		if i==2 {
			goto lco
		}

		if i == 5 {
			// break 

			i++
			continue
		}

		fmt.Println(i)
		i++
	}

	lco: 
	fmt.Println("Jumping to here")
}

