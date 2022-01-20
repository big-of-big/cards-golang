package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	deckSize := 16
	if len(d) != deckSize {
		t.Errorf("Expected deck length of %v, but got %v", deckSize, len(d))
	}

	if d[0] != "スペードのエース" {
		t.Errorf("Expected first card to be 'スペードのエース', but got %v", d[0])
	}

	if d[len(d)-1] != "クローバーの4" {
		t.Errorf("Expected last card to be 'クローバーの4', but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	filename := "_decktesting"
	// もし途中でクラッシュしたらテストで作成したファイルが残っているため
	_ = os.Remove(filename) // _は一時的なファイルであることを意味にする

	deck := newDeck()
	_ = deck.saveToFile(filename)
	loadedDeck := newDeckFromFile(filename)

	deckSize := 16
	if len(loadedDeck) != 16 {
		t.Errorf("Expected %v cards in deck, got %v", deckSize, len(loadedDeck))
	}

	_ = os.Remove(filename)
}
