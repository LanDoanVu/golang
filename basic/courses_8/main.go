package main

import "fmt"

type person struct {
	lastName    string
	firstName   string
	information information
}

type information struct {
	addressHome string
}

func (p person) in() {
	fmt.Printf("%+v", p)
}

func (p *person) changeName(Name string) {
	(*p).firstName = Name
}

func main() {
	// person1 := person{lastName: "Doan Vu", firstName: "Lan"}

	person1 := person{
		lastName:  "Doan Vu",
		firstName: "Lan",
		information: information{
			addressHome: "Ho Chi Minh",
		},
	}

	// person1.lastName = "New"

	// fmt.Printf("Thong Tin %v", person1)

	// fmt.Printf("%#v", person1)

	person1Pointer := &person1
	person1Pointer.changeName("Lin")

	person1.in()
}
