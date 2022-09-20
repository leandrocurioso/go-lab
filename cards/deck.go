package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuites := [4]string{"Hearts", "Diamonds", "Spades", "Clubs"}
	cardValues := [13]string{
		"Ace", "Two", "Three", "Four", "Five", "Six",
		"Seven", "Eight", "Nine", "Ten", "J", "Q", "K",
	}
	for _, cardValue := range cardValues {
		for _, cardSuite := range cardSuites {
			cards = append(cards, cardValue+" of "+cardSuite)
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	content, error := ioutil.ReadFile(filename)
	if error != nil {
		fmt.Println("ERROR WHILE RESTORING A DECK FROM FILE!")
		os.Exit(1)
	}
	s := strings.Split(string(content), ",")
	return deck(s)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
