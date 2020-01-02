package main

import (
	"os"
	"testing"
)

//checking length of cards
func TestNewDeck(t *testing.T) {
	cards := newDeck()
	if len(cards) != 52 {
		t.Errorf("Expected length of 52 but got %v", len(cards))
	}

	if cards[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be Ace of Spades but got %v", cards[0])
	}

	if cards[len(cards)-1] != "King of Clubs" {
		t.Errorf("Expected first card to be King of Clubs but got %v", cards[len(cards)-1])
	}
}

func TestSaveToDeckAndReadDeckFromFile(t *testing.T) {
	os.Remove("_deckTesting")

	deck := newDeck()
	deck.saveToFile("_deckTesting")

	loadedDeck := readDeckFromFile("_deckTesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52 cards but got %v", len(loadedDeck))
	}
	os.Remove("_deckTesting")
}
