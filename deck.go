package main

import "fmt"

// Create a new deck as a slice of strings
type deck []string

// Print the deck
func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

// Create a new deck 
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value + " of " + suit)
		}
	}

	return cards
}