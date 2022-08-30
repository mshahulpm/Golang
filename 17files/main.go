package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fmt.Println("----- Files -----")

	// content := "hey I am a file And I am "
	// file,err :=  os.Create("./hello.txt" )

	// if err != nil {
	// 	panic(err)
	// }

	// length,err := io.WriteString(file,content) 

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("file created and saved content to it length of file : ",length) 

	// defer file.Close()

	readFile("./hello.txt")
}


func readFile(fileName string)  {
 	dataByte,err := ioutil.ReadFile(fileName)

	if err != nil {
		panic(err) 
	}

	fmt.Print("file content : \n",string(dataByte))
}