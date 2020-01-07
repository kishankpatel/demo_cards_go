package main

import "fmt"

type person struct {
	firstName   string
	lastName    string
	contactInfo contactInfo
}

type contactInfo struct {
	email string
	phone string
}

func main() {
	kishan := person{
		firstName: "Kishan",
		lastName:  "Patel",
		contactInfo: contactInfo{
			email: "kishan@gmail.com",
			phone: "7928347934",
		},
	}
	// This updateName function will not update the actual value it will refer to a new address
	kishan.updateName("Kishan Kumar")
	kishan.print()
	
	// getting the address of kishan and updating it, it update the name
	kishanPointer := &kishan
	kishanPointer.update("Kishan Kumar")
	kishan.print()
}

func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func (pk *person) update(newFirstName string) {
	(*pk).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}
