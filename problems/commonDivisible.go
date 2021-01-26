package main

import (
	"fmt"
	"math"
)

func main() {
	a := 5
	b := 10
	n := gcd(a, b)
	p := math.Sqrt(float64(n))
	var limit int = int(p)
	result := 0
	for i := 1; i <= limit; i++ {
		// if 'i' is factor of n
		if n%i == 0 {
			// check if divisors
			// are equal
			if n/i == i {
				result += 1
			} else {
				result += 2
			}
		}
	}
	fmt.Println(result)
}

func gcd(a int, b int) int {
	if a == 0 {
		return b
	}
	return gcd(b%a, a)
}
