package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}
type person struct {
	firstName string
	lastName  string
	//contact   contactInfo   way to use contact var name but not necessary
	contactInfo
}

func main() {

	//alex := person{"Alex", "Anderson"}
	// alex := person{firstName: "Alex", lastName: "Anderson"} //both are ways to initialise structs
	// fmt.Println(alex)

	// var syre person //way to declare alex variable of type person which is a struct
	// syre.firstName = "Rajat"
	// syre.lastName = "Singh"
	// fmt.Println(syre)
	// fmt.Printf("%+v", syre)

	jim := person{
		firstName: "Jim",
		lastName:  "party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 123456,
		},
	}

	// jimPointer := &jim
	// jimPointer.updateName("rajat")
	jim.updateName("jimmy") //pointer shortcut in go, not required upper classic 2 statements for call by references
	jim.print()

}

func (p person) print() { //reciever function for structs
	fmt.Printf("%+v", p)
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}
