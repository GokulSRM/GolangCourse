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
	// gokul := person{firstName: "Gokul", lastName: "Jaeger"}
	// fmt.Println(gokul)

	var gokul person

	// fmt.Println(gokul)

	gokul.firstName = "Gokul"
	gokul.lastName = "Jaeger"
	gokul.contactInfo.email = "gokuljaeger@gmail.com"
	gokul.contactInfo.zipCode = 632002

	sachin := person{
		firstName: "Sachin",
		lastName:  "Hatache",
		contactInfo: contactInfo{
			email:   "sachinhatache20@gmail.com",
			zipCode: 632003,
		},
	}

	// fmt.Printf("%+v", gokul) // %+v will tell you the fields of the struct

	gokul.print()
	fmt.Println("---Before---")
	sachin.print()
	fmt.Println("---After---")
	sachin.updateName("Kakache")
	sachin.print()

	// name := "bill"

	// namePointer := &name

	// fmt.Println(&namePointer)
	// printPointer(namePointer)
}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

// func printPointer(namePointer *string) {
// 	fmt.Println(&namePointer)
// }
