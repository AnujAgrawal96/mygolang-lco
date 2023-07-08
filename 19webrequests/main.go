package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("Web requests demo")

	response, err := http.Get(url)
	checkNilError(err)

	fmt.Printf("Response is of type: %T\n", response)
	defer response.Body.Close()
	//It is the caller's responsibility to close Body.

	databytes, err := ioutil.ReadAll(response.Body)
	checkNilError(err)
	content := string(databytes)
	fmt.Println(content)

}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
