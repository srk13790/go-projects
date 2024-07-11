package main

import "fmt"

type contactInfo struct {
	email string
	zip int
}


// type person struct {
// 	firstName string
// 	lastName  string
// 	contact contactInfo
// }

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	// alex := person{firstName: "Alex", lastName: "Anderson"} //....1

	// var alex person

	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"

	// fmt.Println(alex)
	// fmt.Printf("%+v", alex) //....2

	// jim := person {
	// 	firstName: "Jim",
	// 	lastName: "Party",
	// 	contact: contactInfo{
	// 		email: "jim@email.com",
	// 		zip: 432344,
	// 	},
	// } //....3

	// var jim person

	// jim.firstName = "Jimmy"
	// jim.lastName = "Mayer"
	// jim.contact.email = "Jimm@email.com"
	// jim.contact.zip = 453345 //....4

	jim := person {
		firstName: "Jim",
		lastName: "Party",
		contactInfo: contactInfo{
			email: "jim@email.com",
			zip: 432344,
		},
	} //....5

	
	// fmt.Printf("%+v", jim)

	// jim.print()

	// jimPointer := &jim

	// jimPointer.updateName("Andy")

	// jimPointer.print() ... 6

	jim.print()

	jim.updateName("Andy")

	jim.print()
}


func (p person) print () {
	fmt.Printf("%+v", p)
}


func (pointerToPerson *person) updateName (newFirstName string) {
	(*pointerToPerson).firstName  = newFirstName
}