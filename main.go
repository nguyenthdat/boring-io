package main

import "fmt"

type contactinfo struct {
	email   string
	zipcode int
}

type person struct {
	firstname string
	lastname  string
	contactinfo
}

func main() {
	//alex := person{firstname: "Alex", lastname: "Anderson"}
	//fmt.Println(alex)
	/*
		var alex person
		alex.firstname = "Alex"
		alex.lastname = "Anderson"
		alex.contact.email = "me@alex.com"
		alex.contact.zipcode = 77777
		fmt.Println(alex)
		fmt.Printf("%+v", alex)
	*/
	jim := person{
		firstname: "Jim",
		lastname:  "Party",
		contactinfo: contactinfo{
			email:   "jim@gamil.com",
			zipcode: 9400,
		},
	}
	jim.updatename("Jimmy")
	jim.print()
	name := "bill"
	namePointer := &name
	fmt.Println(&namePointer)
	printPointer(namePointer)
}

func printPointer(namePointer *string) {
	fmt.Println(&namePointer)
}

func (p *person) updatename(newfirstname string) {
	p.firstname = newfirstname
}

func (p person) print() {
	fmt.Println(p)
}
