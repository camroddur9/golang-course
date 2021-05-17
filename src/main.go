package main

import "fmt"

func main() {
	//Constant declaration
	const pi float64 = 3.14
	const pi2 = 3.1415

	fmt.Println(pi, pi2)

	//Variables declaration
	base := 12
	var altura int = 14
	fmt.Println(base * altura)

	// Zero values
	var a int
	var b float64
	var c string
	var d bool
	fmt.Println(a, b, c, d)
}
