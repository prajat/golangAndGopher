package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}
type person struct {
	firstName string
	lastName  string
	contact   contactInfo
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
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 123456,
		},
	}
	fmt.Printf("%+v", jim)

}
