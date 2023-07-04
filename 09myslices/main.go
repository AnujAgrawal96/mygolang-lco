package main

import (
	"fmt"
	"sort"
)

func main() {
	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("Type of fruitList is %T \n", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList)

	fruitList = append(fruitList[:3])
	fmt.Println(fruitList)

	highscores := make([]int, 4)

	highscores[0] = 234
	highscores[1] = 945
	highscores[2] = 465
	highscores[3] = 867
	//highscores[4] = 777
	//out of bound undex error

	fmt.Println("address of highscores : ", &highscores[0])

	//append reallocates memory
	highscores = append(highscores, 777, 888)
	fmt.Println("address of highscores : ", &highscores[0])

	fmt.Println(highscores)

	//sorting slices
	fmt.Println(sort.IntsAreSorted(highscores))
	sort.Ints(highscores)
	fmt.Println("sorted slice : ", highscores)

	//remove a value from slices based on index
	var courses = []string{"react", "js", "swift", "python", "java"}
	fmt.Println(courses)
	var index = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}
