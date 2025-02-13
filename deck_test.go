package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := NewDeck()

	if len(d) != 56 {
		t.Errorf("Expcted deck length of 16, but got %v", len(d))
	}
}