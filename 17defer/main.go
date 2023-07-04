package main

import "fmt"

func main() {
	//example given ni go doc using lock() defer unlock()
	//can only be understood after mutex, sync and memory shared functions

	//1. defer functions are invoked immediately before the surrounding function returns
	//2. invoked in the LIFO order (reverse order of being deferred)

	defer fmt.Println("World")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("Hello")

	myDefer()

	fmt.Println("")
}

//world, one, two
//0, 1, 2, 3, 4
//hello, 43210, two, one, world

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}
