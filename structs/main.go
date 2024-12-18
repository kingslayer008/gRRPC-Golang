package main

import "fmt"

type contactInfo struct {
	email   string
	zipcode int
}

type person struct {
	firstname string
	lastname  string
	contact   contactInfo
}

func (pointerToPerson *person) updateName(firstname string) {
	(*pointerToPerson).firstname = firstname
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func main() {
	jim := person{
		firstname: "Dineshkanna",
		lastname:  "M",
		contact: contactInfo{
			email:   "dinesh21a@gmail.com",
			zipcode: 638181,
		},
	}
	jim.updateName("Dinesh")
	jim.print()

}
