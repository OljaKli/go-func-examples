package main

import "fmt"

type person struct {
	first string
	last  string
}

type administrator struct {
	person
	admin bool
}

func (a administrator) Login() {
	fmt.Println("User", a.first, a.last, "is login.")
}

func methods() {
	admin1 := administrator{
		person: person{
			first: "John",
			last:  "Doe",
		},
		admin: true,
	}
	admin1.Login()
}
