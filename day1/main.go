package main

import "fmt"

func main() {
	//function add ini adalah logika pertambahan yang ada di calculator.go yang mempunyai 2 parameter yang dimana bertipe integer
	// dan memiliki return value integer jadi function ini mengembalikan parameter a + b
	fmt.Println(add(10, 5))

	//function substract ini adalah logika pengurangan yang ada di calculator.go
	// mempunyai 2 parameter yang dimana bertipe integer
	// dan memiliki return value integer jadi function ini mengembalikan parameter a - b
	fmt.Println(substract(10, 5))

	//function multiply ini adalah logika perkalian yang ada di calculator.go
	// mempunyai 2 parameter yang dimana bertipe integer
	// dan memiliki return value integer jadi function ini mengembalikan parameter a * b
	fmt.Println(multiply(10, 5))

	//function divide ini adalah logika pembagian yang ada di calculator.go
	// mempunyai 2 parameter yang dimana bertipe integer
	// dan memiliki return value integer jadi function ini mengembalikan parameter a / b
	fmt.Println(divide(10, 5))
}
