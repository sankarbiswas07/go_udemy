package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 20, but got %v", len(d))
	}
	
	if d[1] != "Two of Spades" {
		t.Errorf("Expected second card as Two of Spades, but got %v", d[1])
	}

	if d[len(d) - 1] != "King of Clubs" {
		t.Errorf("Expected last card as King of Clubs, but got %v", d[1])
	}

}