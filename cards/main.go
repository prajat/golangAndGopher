package main

import "fmt"

func main() {
	//var card string = "Ace of spades"
	card := newCard()
	//card = "five of hearts"
	fmt.Println(card)
}

func newCard() string {
	return "four of hearts"
}
