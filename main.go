package main

import (
	"fmt"
)

func main() {
	cards := deck{newcard(), newcard()}
	cards = append(cards, "HAHAHAHAHA")
	cards.print()
	fmt.Println([]byte("Hi There"))

	greeding := "Hi Jason"
	fmt.Println([]byte(greeding))

}

func newcard() string {
	return "Trible fast"
}
