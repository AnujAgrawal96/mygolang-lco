package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=hkjhskjhjhajsdg"

func main() {
	fmt.Println("Welcome to handling URLs in golang")
	fmt.Println(myurl)

	//parsing the url
	response, err := url.Parse(myurl)
	checkNilError(err)
	fmt.Printf("Response is of type: %T\n", response)

	fmt.Println(response.Scheme)
	fmt.Println(response.Host)
	fmt.Println(response.Path)
	fmt.Println(response.Port())
	fmt.Println(response.RawQuery)

	qparams := response.Query()
	fmt.Printf("Query params is of type: %T\n", qparams)
	fmt.Println(qparams)

	fmt.Println(qparams["coursename"])

	for key, value := range qparams {
		fmt.Printf("Param key is %v and value is %v\n", key, value)
	}

	//constructing a url
	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "tutcss",
		RawPath: "user=anuj",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
