package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
)

type Deck []string
func NewDeck() Deck {
	var newDeck Deck

	cardSuits := []string {"Hearts", "Diamonds", "Clubs", "Spades"}
	cardValues := []string {"Ace", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

	for _, suite := range cardSuits {
		for _, value := range cardValues {
			newDeck = append(newDeck, value + " of " + suite)
		}
	}

	return newDeck
}

func (d Deck) Print() {
	for _, card := range d {
		fmt.Print(card + "|")
	}
}

func Deal(d Deck, handSize int) (Deck, Deck) {
	return d[:handSize], d[handSize:]
}

func (d Deck) toString() string {
	return strings.Join([]string(d), ",")
}

func toSlice(joinedDeck string) []string {
	return strings.Split(joinedDeck, ",")
}

func (d Deck) saveToFile(filename string) error {
	data := []byte(d.toString())

	err := os.WriteFile(filename, data, 0666)

	if err != nil {
		log.Fatal(err)
	}

	return err
}

func newDeckFromFile(filename string) Deck {
	data, err := os.ReadFile(filename)

	if err != nil {
		log.Fatal(err)
	}

	deckString := string(data)
	deckSlice := toSlice(deckString)

	return Deck(deckSlice)
}

func (d Deck) Shuffle() Deck {
	rand.Shuffle(len(d), func(i, j int) {
		d[i], d[j] = d[j], d[i]
	})

	return d
}
