package main

func main() {
	cards := deck{"Ace of shit", newCard()}
	cards = append(cards, "six of mowzes")
	cards.print()
}

func newCard() string {
	return "peace of shits"
}
