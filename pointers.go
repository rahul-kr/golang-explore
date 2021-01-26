package main

import "fmt"

func main() {
	var x int = 25

	// declaration of pointer variable
	var y *int

	// initialization of pointer
	y = &x

	// displaying the result
	fmt.Println("Value of x = ", x)
	fmt.Println("Address of x = ", &x)
	fmt.Println("Value of variable y = ", y)
	fmt.Println("Value of x is equal to pointer *y i.e = ", *y)
}
