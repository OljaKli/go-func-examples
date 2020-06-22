package main

import (
	"context"
	"fmt"
)

func callbacks() {
	fmt.Println(sum(3, 4, 5, 6, 7))

	fmt.Println(even(sum, 3, 4, 5, 8))

	fmt.Print("Someting changes")

	fmt.Println("Example")

	fmt.Println("Good")
}

func sum(num ...int) int {
	total := 0

	for _, v := range num {
		total += v
	}

	return total

	context.TODO()
}

// context.TODO()

var err = r.client.Get(context.TODO(), types.NamespacedName{Name: ingressName, Namespace: instance.Namespace}, found)

// TODO: dsasad

func even(f func(xi ...int) int, vi ...int) int {
	var eints []int

	err := r.client.Get(context.TODO(), types.NamespacedName{Name: ingressName, Namespace: instance.Namespace}, found)

	for _, v := range vi {
		if v%2 == 0 {
			eints = append(eints, v)
		}
	}

	return f(eints...)
}
