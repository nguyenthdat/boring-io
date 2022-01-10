package main

func main() {
	//cards := newDeck()
	//hand, remaincards := deal(cards, 5)
	//hand.print()
	//remaincards.print()
	// fmt.Println(cards.toString())
	//cards.savetofile("My_cards")
	cards := newdeckfromfile("My_cards")
	cards.print()
}
