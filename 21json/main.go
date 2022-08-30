package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("------JSON------")
//    encodeJson()
decodeJson()
}


type course struct {
	Name     string `json:"course_name"`
	Price    int 
	Platform string 
	Password string `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}


func encodeJson(){

	courses := []course{
		{"React",299,"hello.com","12345gter",[]string{"web","dev","react"}},
		{"Node",299,"hello2.com","hayt45gter",[]string{"web","dev","react"}},
		{"React",299,"hello.com","12345gter",nil},
	}

	finalJson,err := json.MarshalIndent(courses,"","\t")

	if err != nil {
		panic(err)
	}


   fmt.Printf("%s\n",finalJson)

}


func decodeJson(){
	jsonDataFromWeb := []byte(`
	{
		"course_name": "Node",
		"Price": 299,
		"Platform": "hello2.com",
		"tags": ["web","dev","react"]
}
	`)

	var _course course 
	
	checkValid := json.Valid(jsonDataFromWeb) 

	if checkValid {
		fmt.Println("JSON was valid") 
		json.Unmarshal(jsonDataFromWeb,&_course) 
		fmt.Printf("%#v\n",_course)
	}else {
		fmt.Println("Invalid json")
	}

	// sometime we need data as key value 

	var myOnlineData map[string]interface{} 
	json.Unmarshal(jsonDataFromWeb,&myOnlineData) 
	fmt.Printf("%#v\n",myOnlineData)

	for k,v := range myOnlineData {
		fmt.Printf("Key is %v and value %v and type %T\n",k,v,v)
	}
}