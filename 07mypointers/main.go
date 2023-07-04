package main

import "fmt"

func main() {

	//var ptr *int
	//fmt.Println("Pointer value ", ptr)

	//o/p : Pointer value  <nil>

	myNumber := 23
	var ptr = &myNumber
	fmt.Println("ptr actual value : ", ptr)
	fmt.Println("myNumber address  : ", &myNumber)
	fmt.Println("ptr referenced value : ", *ptr)
	fmt.Println("ptr address: ", &ptr)

	*ptr = *ptr + 2
	fmt.Println("myNumber : ", myNumber)

}
