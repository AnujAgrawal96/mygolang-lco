package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("welcome to web verb.")
	PerformGetRequest()
	performPostJsonRequest()
	performPostFormRequest()
}

func PerformGetRequest() {
	const myurl = "http://localhost:1111/get"

	response, err := http.Get(myurl)
	checkNilError(err)

	defer response.Body.Close()

	fmt.Println("Status Code :", response.StatusCode)
	fmt.Println("Content length is:", response.ContentLength)

	//approach 1:
	/* 	content, err := ioutil.ReadAll(response.Body)
	   	checkNilError(err)

	   	fmt.Println(string(content)) */

	//aproach 2: using builder library
	var responseString strings.Builder
	content, err := ioutil.ReadAll(response.Body)
	checkNilError(err)
	byteCount, _ := responseString.Write(content)

	fmt.Println("ByteCount is:", byteCount)

	fmt.Println(string(content))
}

func performPostJsonRequest() {
	const myurl = "http://localhost:1111/post"

	//fake json payload
	requestBody := strings.NewReader(`
		{
			"coursename":"json post request in golang",
			"price":0,
			"platform":"lco.in"
		}
	`)

	response, err := http.Post(myurl, "application/json", requestBody)
	checkNilError(err)

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}

func performPostFormRequest() {
	const myurl = "http://localhost:1111/postform"

	//formdata
	data := url.Values{}
	data.Add("firstname", "anuj")
	data.Add("lastname", "agrawal")
	data.Add("email", "anuj@go.dev")

	response, err := http.PostForm(myurl, data)
	checkNilError(err)

	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
