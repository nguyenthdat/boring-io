package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Expecting deck has 16 cards, but got %v cards", len(d))
	}
	if d[0] != "Ace of Spades" {
		t.Errorf("Expecting first card is Ace of Spades, but got %v", d[0])
	}
	if d[len(d)-1] != "Four of Rings" {
		t.Errorf("Expecting last card is Four of Rings, but got %v", d[len(d)-1])
	}
}

func TestIO(t *testing.T) {
	os.Remove("_decktesting")
	deck := newDeck()
	deck.savetofile("_decktesting")
	loaddeck := newdeckfromfile("_decktesting")
	if len(loaddeck) != 16 {
		t.Errorf("Expecting deck has 16 cards, but got %v cards", len(loaddeck))
	}
	os.Remove("_decktesting")
}
