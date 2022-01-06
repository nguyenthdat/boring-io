package main

import "fmt"

func main() {
	cards := []string{newcard(), newcard()}
	cards = append(cards, "HAHAHAHAHA")
	for i := 0; i < 3; i++ {
		fmt.Println(cards[i])
	}

	for _, card := range cards {
		fmt.Println(card)
	}
}

func newcard() string {
	return "Trible fast"
}
