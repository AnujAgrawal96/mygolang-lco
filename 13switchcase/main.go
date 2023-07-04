package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Number on dice is :", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("dice value is 1 and you can open")
	case 2:
		fmt.Println("move 2 spot")
	case 3:
		fmt.Println("move 3 spot")
		fallthrough
	case 4:
		fmt.Println("move 4 spot")
		fallthrough
	case 5:
		fmt.Println("move 5 spot")
	case 6:
		fmt.Println("move 6 spot and roll dice again")
	default:
		fmt.Println("wrong input!!")
	}
}
