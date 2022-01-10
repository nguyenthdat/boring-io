package main

import "fmt"

func main() {
	cards := newDeck()
	//hand, remaincards := deal(cards, 5)
	//hand.print()
	//remaincards.print()
	fmt.Println(cards.toString())
}
