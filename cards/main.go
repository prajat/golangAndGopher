package main

func main() {
	cards := newDeck()
	hand, remainingCards := deal(cards, 5)
	hand.print()
	remainingCards.print()

	//cards.print()

}

// func newCard() string {
// 	return "four of hearts"
// }

// variables	//var card string = "Ace of spades"
// card := newCard()
//card = "five of hearts"

//arrays and slices and loops
// cards := []string{"ace of diamonds", newCard()}
// cards = append(cards, "three of spades")
// for i, card := range cards {
// 	fmt.Println(i, card)
// }
