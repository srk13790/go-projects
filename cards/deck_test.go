package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d:= newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	if d[0] != "Ace Of Spades" {
		t.Errorf("Expected first card as Ace of Spades, But got %v", d[0])
	}

	if d[len(d) - 1] != "Four Of Clubs" {
		t.Errorf("Expected first card as Four of Clubs, But got %v", d[len(d) - 1])
	}
}

func TestSaveToDeckAndNewDeckFromFile (t *testing.T) {
	os.Remove("_docktesting")

	deck := newDeck()
	deck.saveToFile("_docktesting")

	loadedDeck := newDeckFromFile("_docktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(loadedDeck))
	}

	os.Remove("_docktesting")
}