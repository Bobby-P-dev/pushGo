package main

import "fmt"

func main() {

	user := User{
		Name: "Bobby",
		Age:  20,
	}

	fmt.Println(user.GetName())

	user.ChangeName("Alice")

	fmt.Println(user.GetName())
}
