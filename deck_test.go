package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Expecting deck has 16 cards, but got %v cards", len(d))
	}
}
