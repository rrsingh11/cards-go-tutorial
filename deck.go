package main

import (
	"fmt"
	"os"
	"strings"
	"time"
	"math/rand"
)

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

// Convert []string to a single string sep by ","
func (d deck) toString() string{
	return strings.Join([]string(d), ",")
} 

// Save to File
func (d deck) saveToFile(name string) {
	err := os.WriteFile(name, []byte(d.toString()), 0666)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}	
}

// Get deck from file
func getDeckFromFile(file string) deck {
	bs, err := os.ReadFile(file)
	if(err != nil) {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	return deck(strings.Split(string(bs), ","))
}

// Shuffle deck
func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano()) // Each time convert it into int64
	r := rand.New(source) // create new seed

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}

// Create a hand
func deal(d deck, num int) (deck, deck) {
	return d[:num], d[num:]
}