package main

import (
	"fmt"
	"math/big"

	//"math/rand"
	"crypto/rand"
)

func main() {
	//var mynumberOne int = 2
	//var mynumberTwo float64 = 4.5

	//precision would be lost
	//fmt.Println("Sum :", mynumberOne+int(mynumberTwo))

	//fmt.Println("Sum :", strconv.)

	//random number
	//rand.Seed(time.Now().UnixNano())
	//fmt.Println(rand.Intn(5))

	//random from crypto
	myRandomNumber, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(myRandomNumber)

}
