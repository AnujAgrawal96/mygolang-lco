package main

import "fmt"

func main() {

	//there are immediately executed functions,lambda functions , annonymous functions ,

	result, status := sum(3, 5)
	fmt.Println("Sum:", result)
	fmt.Println("Status :", status)

	proResult, proStatus := proSum(2, 5, 8, 7)
	fmt.Println("Pro Result :", proResult)
	fmt.Println("Pro Status :", proStatus)
}

func proSum(nums ...int) (int, string) {
	total := 0

	for _, num := range nums {
		total += num
	}

	return total, "success"
}

func sum(num1 int, num2 int) (int, string) {
	return num1 + num2, "success"
}
