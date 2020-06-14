package main

import "fmt"

func callbacks() {
	fmt.Println(sum(3, 4, 5, 6, 7))

	fmt.Println(even(sum, 3, 4, 5, 8))
}

func sum(num ...int) int {
	total := 0

	for _, v := range num {
		total += v
	}

	return total
}

func even(f func(xi ...int) int, vi ...int) int {
	var eints []int

	for _, v := range vi {
		if v%2 == 0 {
			eints = append(eints, v)
		}
	}

	return f(eints...)
}
