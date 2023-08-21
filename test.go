package main

import "fmt"

func test() {
	cards := []string{"Ace", "2", "3"}

	for _, card := range cards{
		fmt.Println(card)
	}
}