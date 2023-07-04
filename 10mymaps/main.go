package main

import "fmt"

func main() {
	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["JV"] = "Java"
	languages["PY"] = "Python"

	fmt.Println("map : ", languages)
	fmt.Println("JS shorts for :", languages["JS"])

	delete(languages, "JV")
	fmt.Println("map : ", languages)

	//looping map
	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}

}
