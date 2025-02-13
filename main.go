package main

func main() {
	cards := NewDeck()
	cards.Shuffle()
	cards.saveToFile("saved.txt")
	newDeckFromFile("saved.txt")
	cards.Print()
}