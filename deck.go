package main

import "fmt"

type deck []string

// reciever function
func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}
