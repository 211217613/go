package main

import "fmt"

func main() {
	// slice
	// cards := newDeck()
	// cards.saveToFile("my_cards.txt")

	cards := newDeck()
	cards.print()
	fmt.Println()
	cards.shuffle()
	cards.print()
}
