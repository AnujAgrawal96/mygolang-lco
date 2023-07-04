package main

import "fmt"

func main() {
	fmt.Println("structs in golang")
	// no inheritance , no super , no child , no parent

	anuj := User{"Anuj", "anuj@go.dev", true, 27}
	fmt.Println(anuj)

	//+v used to print structure in readable format
	fmt.Printf("anuj details are: %+v\n", anuj)
	fmt.Printf("Name is %v and email is %v.\n", anuj.Name, anuj.Email)
	anuj.GetStatus()

	anuj.NewMmail()
	//notice email is not changed in actual user object
	fmt.Printf("anuj details are: %+v\n", anuj)
}

type User struct {
	//capital public(needs to be exported) , small private
	Name   string
	Email  string
	Status bool
	Age    int
}

//defining methods of struct like class
func (u User) GetStatus() {
	fmt.Println("Is User active :", u.Status)
}

//copy is passed thats why pointers are designed
func (u User) NewMmail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is:", u.Email)
}
