package main

func main() {
	cards := deck{"Ace of Diamond", newcard()}
	cards = append(cards, "Six of Spades")
	cards.print()
}

func newcard() string {
	return "Five of Diamonds"
}
