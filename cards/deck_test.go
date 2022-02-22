package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 28 {
		t.Errorf("Expected deck length of 27 but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}
}

func TestSaveToDeskAndNewDeckFromFile(t *testing.T){
	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")
	loadedDeck := newDeckFromFile("_decktesting")
	if len(loadedDeck) != 28 {
		t.Errorf("Expected 28 cards in the deck, got %v", len(loadedDeck))
	}
}
