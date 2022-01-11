package main

import "fmt"

type contactinfo struct {
	email   string
	zipcode int
}

type person struct {
	firstname string
	lastname  string
	contact   contactinfo
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
		contact: contactinfo{
			email:   "jim@gamil.com",
			zipcode: 9400,
		},
	}
	fmt.Println(jim)
}
