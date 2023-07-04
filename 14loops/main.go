package main

import "fmt"

func main() {
	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}
	fmt.Println(days)

	fmt.Println("using for loop")
	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}
	fmt.Println("")

	fmt.Println("Using for range")
	for i := range days {
		fmt.Println(days[i])
	}
	fmt.Println("")

	fmt.Println("using for each loop")
	for index, day := range days {
		fmt.Printf("index is %v and value is %v\n", index, day)
	}

	fmt.Println("break and continue")

	rougeValue := 1

	for rougeValue < 10 {

		if rougeValue == 2 {
			goto stmt1
		}

		if rougeValue == 8 {
			rougeValue++
			break
		}

		if rougeValue == 5 {
			rougeValue++
			continue
		}

		fmt.Println("value is :", rougeValue)
		rougeValue++
	}

	//goto statements

stmt1:
	fmt.Println("jumped to stmt1")

}
