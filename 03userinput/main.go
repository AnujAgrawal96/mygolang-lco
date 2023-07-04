package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	greetings := "Welcome to possible!!"
	fmt.Println(greetings)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our Pizza:")

	//comma ok || comma err syntax

	//read till '\n'
	input, _ := reader.ReadString('\n')

	//input, err := reader.ReadString('\n')
	//err for error handling

	fmt.Println("Thanks for rating us :", input)
	fmt.Printf("Type of rating is %T", input)

}
