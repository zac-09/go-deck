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
	// alex := person{firstName: "Alex", lastName: "Anderson"}
	// fmt.Println(alex)
	// var alex person
	// fmt.Printf("struct %+v", alex)
	jim := person{
		firstName: "jim",
		lastName:  "Bisnga",
		contactInfo: contactInfo{
			email:   "isa@gmai.com",
			zipCode: 847541,
		},
	}

	jim.updateName("isaac")
	jim.print()

}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}
func (p person) print() {
	fmt.Printf("%+v", p)

}
