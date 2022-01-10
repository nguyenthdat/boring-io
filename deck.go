package main

import "fmt"

type deck []string

// reciever function
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
