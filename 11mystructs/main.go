package main

import "fmt"

func main() {
	fmt.Println("structs in golang")
	// no inheritance , no super , no child , no parent

	anuj := User{"Anuj", "anuj@go.dev", true, 27}
	fmt.Println(anuj)

	//+v used to print structure in readable format
	fmt.Printf("anuj details are: %+v\n", anuj)
	fmt.Printf("Name is %v and email is %v.", anuj.Name, anuj.Email)
}

type User struct {
	//capital public(needs to be exported) , small private
	Name   string
	Email  string
	Status bool
	Age    int
}
