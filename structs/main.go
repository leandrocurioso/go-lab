package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode string
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	person1 := person{
		firstName: "Leandro",
		lastName:  "Curioso",
		contact: contactInfo{
			email:   "leandro.curioso@gmail.com",
			zipCode: "04122000",
		},
	}

	person1.updateName("NutzUnkind")
	person1.print()
}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
