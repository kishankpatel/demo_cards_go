package main

import "fmt"

func main() {

	// for _, card := range cards {
	// 	fmt.Println(card)
	// }
	// cards := newDeck()
	// hand, remainingDeck := deal(cards, 6)
	// hand.print()
	// fmt.Println("--")
	// remainingDeck.print()

	cards := newDeck()
	cards.saveToFile("CardList.txt")

	loadedCards := readDeckFromFile("CardList.txt")

	fmt.Println("Before suffeling:")
	loadedCards.print()

	fmt.Print("\nSuffeling...\n\n")
	loadedCards.suffle()

	fmt.Println("After suffeling:")
	loadedCards.print()
}

func newCard() string {
	return "Ace of Spades"
}
