package main

import "fmt"

func main() {
	loginCount := 23
	var result string

	if loginCount < 10 {
		result = "Regular User"
	} else if loginCount > 10 {
		result = "frequent user"
	} else {
		result = "exactly 10 count"
	}

	fmt.Println(result)

	// on the go condition check

	if 9%2 == 0 {
		fmt.Println("even number")
	} else {
		fmt.Println("odd number")
	}

	// web request handling

	if num := 3; num < 10 {
		fmt.Println("number is less than 10")
	} else {
		fmt.Println("number is not less than 10")
	}

	// if err != nil {
	// 	return nil, err
	// }

}
