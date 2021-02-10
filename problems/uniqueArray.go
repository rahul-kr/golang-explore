package main

import (
	"fmt"
	"sort"
)

func main() {
	randomValues := []int{5, 5, 6, 8, 4, 7, 6}
	var uniqueValues []int
	sort.Ints(randomValues)
	var matchValue int
	for index, value := range randomValues {
		if matchValue != value {
			uniqueValues = append(uniqueValues, value)
		}
		matchValue = value

		fmt.Printf("Index is: %d -- Value is: %d\n", index, value)
	}
	fmt.Println(uniqueValues)
}
