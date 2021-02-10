package main

import "fmt"

func main() {
	var stack []int
	push(5, &stack)
	push(6, &stack)
	push(7, &stack)
	push(8, &stack)
	fmt.Println(stack)
	if len(stack) != 0 {
		pop(&stack)
		fmt.Println(stack)
	}
	if len(stack) != 0 {
		pop(&stack)
		fmt.Println(stack)
	}
	if len(stack) != 0 {
		pop(&stack)
		fmt.Println(stack)
	}
	if len(stack) != 0 {
		pop(&stack)
		fmt.Println(stack)
	}
	if len(stack) != 0 {
		pop(&stack)
		fmt.Println(stack)
	}
}

func push(val int, stack *[]int) {
	*stack = append(*stack, val)
}

func pop(stack *[]int) {
	stackLenght := len(*stack) - 1
	*stack = (*stack)[:stackLenght]
}
