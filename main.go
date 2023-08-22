package main

// import "fmt"

func main() {
	cards := newDeck()
	// cards := getDeckFromFile()
	cards.shuffle()

	// hand, remainingCards := deal(cards, 5)

	cards.print()
	// hand.print()
	// cards.saveToFile()
}