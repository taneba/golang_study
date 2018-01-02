package main

func main() {
	// cards := newDeckFromFile("my")
	// cards.print()
	cards := newDeck()
	cards.shuffle()
	cards.print()
}
