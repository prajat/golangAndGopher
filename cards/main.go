package main

import "fmt"

func main() {

	// variables

	//var card string = "Ace of spades"
	// card := newCard()
	//card = "five of hearts"
	// fmt.Println(card)

	//arrays and slices

	cards := []string{"ace of diamonds", newCard()}
	cards = append(cards, "three of spades")

	for i, card := range cards {
		fmt.Println(i, card)
	}
	//fmt.Println(cards)

}

func newCard() string {
	return "four of hearts"
}
