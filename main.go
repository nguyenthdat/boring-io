package main

func main() {
	cards := deck{newcard(), newcard()}
	cards = append(cards, "HAHAHAHAHA")
	cards.print()
}

func newcard() string {
	return "Trible fast"
}
