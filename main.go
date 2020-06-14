package main

import "fmt"

func main() {
	methods()
	interfaces()

	func(age int) {
		fmt.Println("Anonymous function: age", age)
	}(22)

	expression()
}
