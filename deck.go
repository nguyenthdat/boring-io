package main

import "fmt"

type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuites := []string{"Spades", "Diamonds", "Hearts", "Rings"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}
	for _, suite := range cardSuites {
		for _, value := range cardValues {
			cards = append(cards, suite+" of "+value)
		}
	}
	return cards
}

// reciever function
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handsize int) (deck, deck) {
	return d[:handsize], d[handsize:]
}
