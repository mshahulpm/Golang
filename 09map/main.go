package main

import "fmt"

func main() {
	fmt.Println("Maps")

	languages := make(map[string]string) 

	languages["js"] = "javascript"
	languages["ts"] = "typescript"
	languages["py"] = "python"
	languages["rb"] = "ruby"

	fmt.Println("list of all languages ",languages)
    fmt.Println("js short ",languages["js"])

	delete(languages,"rb")
	fmt.Println(languages)

	// looping through map 
	for key,value := range languages {
		fmt.Printf("For key %v, value %v\n",key,value) 
	}
}