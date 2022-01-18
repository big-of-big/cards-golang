package main

import "testing"

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
