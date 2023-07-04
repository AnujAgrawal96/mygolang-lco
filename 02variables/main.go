package main

import "fmt"

var numberOfUser2 int = 3000

//numberOfUser2 := 3000 (outside method not allowed)

const LoginToken string = "jdsjdkjahjh238283" //declared public as starts from capital letter

func main() {
	fmt.Println("Variables")

	var username string = "Anuj"
	fmt.Println(username)
	fmt.Printf("Variables of type: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variables of type: %T \n", isLoggedIn)

	//var smallval uint8 = 256
	var smallval uint8 = 255
	fmt.Println(smallval)
	fmt.Printf("Variables of type: %T \n", smallval)

	var smallFloat float32 = 255.455361628781272
	fmt.Println(smallFloat)
	fmt.Printf("Variables of type: %T \n", smallFloat)

	//default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variables of type: %T \n", anotherVariable)

	var anotherVariable1 string
	fmt.Println(anotherVariable1)
	fmt.Printf("Variables of type: %T \n", anotherVariable1)

	var anotherVariable2 float64
	fmt.Println(anotherVariable2)
	fmt.Printf("Variables of type: %T \n", anotherVariable2)

	//implicit type
	var website = "anuj.com"
	fmt.Println(website)
	//website = 3
	// will throw IncompatibleAssign error as already string type

	//no var style
	numberOfUser := 3000
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("Variables of type: %T \n", LoginToken)
}
