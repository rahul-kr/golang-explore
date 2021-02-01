package main

import (
	"fmt"
)

func main() {
	a := 5
	defer fmt.Println("value of a is ", a)
	a += 1
	fmt.Println("value of a is ", a)

}
