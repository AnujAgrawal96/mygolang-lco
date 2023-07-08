package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files in Golang")

	content := "First line of a file"

	file, error := os.Create("./mygofile.txt")

	/* if error != nil {
		panic(error)
	} */
	checkNilError(error)

	length, err := io.WriteString(file, content)
	checkNilError(err)

	fmt.Println("length is:", length)
	defer file.Close()

	readFile("./mygofile.txt")
}

func readFile(filename string) {
	dataBytes, err := ioutil.ReadFile(filename)
	checkNilError(err)

	fmt.Println("Text data inside the file is \n", dataBytes)

	fmt.Println("Text data inside the file is \n", string(dataBytes))

}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
