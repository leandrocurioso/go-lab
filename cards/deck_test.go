package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	dc := 52
	l := len(d)
	if l != dc {
		t.Errorf("Expected deck length of %v but got %v", dc, l)
	}

	c1 := "Ace of Hearts"
	fc1 := d[0]
	if fc1 != c1 {
		t.Errorf("Expected first card to be %v but got %v", c1, fc1)
	}

	c2 := "K of Clubs"
	fc2 := d[len(d)-1]
	if fc2 != c2 {
		t.Errorf("Expected last card to be %v but got %v", c2, fc2)
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	dtf := "_decktesting"
	os.Remove(dtf)
	d := newDeck()
	d.saveToFile(dtf)
	dc := 52
	ld := newDeckFromFile(dtf)
	cld := len(ld)
	if cld != dc {
		t.Errorf("Expected %v total of cards but got %v", dc, cld)
	}
	os.Remove(dtf)
}
