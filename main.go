package main

import "fmt"

func main() {
	cards := deck{"Ace of shit", newCard()}
	cards = append(cards, "six of mowzes")
	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "peace of shits"
}
