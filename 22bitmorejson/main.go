package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to JSON!!")
	//EncodeJson()
	DecodeJson()
}

func EncodeJson() {

	courses := []course{
		{"ReactJs Bootcamp", 299, "w3schools.in", "abc123", []string{"web-dev", "js"}},
		{"MERN Bootcamp", 199, "w3schools.in", "bcd123", []string{"full-stack", "js"}},
		{"Angular Bootcamp", 299, "w3schools.in", "hit123", nil},
	}

	//package this data as json data
	//finalJson, err := json.Marshal(courses)
	finalJson, err := json.MarshalIndent(courses, "", "\t")
	checkNilError(err)

	fmt.Printf("%s\n", finalJson)

}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
		{
			"coursename": "ReactJs Bootcamp",
			"Price": 299,
			"website": "w3schools.in",  
			"tags": ["web-dev","js"]
	    }
	`)

	var myCourse course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("json data is valid!!")
		json.Unmarshal(jsonDataFromWeb, &myCourse)

		// %#v - for interfaces/structs
		fmt.Printf("%#v\n", myCourse)
	} else {
		fmt.Println("json was not valid")
	}

	//cases where we just want to add data to key value pair

	//an interface type is defined as a set of method signatures
	//type Abser interface{
	//Abs() float64
	//}

	fmt.Println("")
	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	// %#v - for interfaces/structs
	fmt.Printf("%#v\n", myOnlineData)

	fmt.Println("")
	for key, value := range myOnlineData {
		fmt.Printf("Key is %v and value is %vand type is %T\n", key, value, value)
	}

}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
