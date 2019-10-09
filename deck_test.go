package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Expected desk length of 16,but got %v", len(d))
	}

	if d[0] != "ace of spades" {
		t.Errorf("Expected first card ace of spades,but got %v", d[0])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")
	deck := newDeck()
	deck.saveToFile("_decktesting")
	loadedDeck := newDeckFromFile("_decktesting")
	if len(loadedDeck) != 16 {
		t.Errorf("Expected desk length of 16,but got %v", len(loadedDeck))
	}
	os.Remove("_decktesting")
}
