package main

import "fmt"

type manager interface {
	Login()
}

func Delete(m manager) {
	fmt.Println(m, "user delete")
}

func interfaces() {
	admin2 := administrator{
		person: person{
			first: "James",
			last:  "Smite",
		},
		admin: true,
	}
	Delete(admin2)
}
