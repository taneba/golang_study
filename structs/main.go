package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	// こういう書き方もできるが
	// alex := person{"Alex", "Anderson"}
	// struct定義時の順番が変わると意図しないプロパティに値が入ってしまうので良くない
	// alex := person{firstName: "Alex", lastName: "Anderson"}
	// fmt.Println(alex)
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}
	// jimPointer := &jim

	// &jim : give me access to the memory address that this variable is sitting at.
	jim.updateName("jimmy")
	jim.print()
}

// this update main function can only be called with the receiver of a pointer to a person
func (pointerToPerson *person) updateName(newFirstName string) {
	// *person : a type of pointer that points at a pserson
	(*pointerToPerson).firstName = newFirstName
	// *pointerToPerson : give me the value this memory address is pointin at
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
