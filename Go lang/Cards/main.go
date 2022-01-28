package main

import "fmt"

func main() {

	cards := newDeck()

	// fmt.Println("Complete deck:")
	// cards.printDeck()

	cards.shuffle()
	fmt.Println("Shuffled deck:")
	cards.printDeck()

	// cards.saveToFile("All cards")
	// fmt.Println(readDeckFromFile("All cards"))

	// hand, remaningCards := deal(cards, 5)

	// fmt.Println("\nHand:")
	// hand.printDeck()

	// fmt.Println("\nRemaining cards:")
	// remaningCards.printDeck()
}
