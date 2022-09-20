package main

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()
	/*
		filenameStateFile := "C:\\Users\\leand\\Projects\\go-lab\\cards\\deck-dump.state"
		hand, _ := deal(cards, 10)
		hand.saveToFile(filenameStateFile)
		restoredDeck := newDeckFromFile(filenameStateFile)
		fmt.Println(restoredDeck)
	*/
}
