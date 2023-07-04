package main

import "fmt"

func main() {
	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[3] = "Peach"

	fmt.Printf("Type of fruitList is : %T \n ", fruitList)

	fmt.Println("ArrayList : ", fruitList)
	fmt.Println("length of ArrayList : ", len(fruitList))

	var vegList = [5]string{"potato", "beans", "mushroom"}
	fmt.Println("vegList : ", vegList)
}
