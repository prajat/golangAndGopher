package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}

type bot interface {
	getGreeting() string
}

func main() {

	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)

}

func (englishBot) getGreeting() string {
	//very custom language for english
	return "hi there!"
}

func (spanishBot) getGreeting() string {
	return "hola!"
}
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// func printGretting(eb englishBot) {
// 	fmt.Println(eb.getGreeting())
// }                                                 //dont need to make 2 function of same logic instead weel use interface and make only one function

// func printGretting(sb spanishBot) {
// 	fmt.Println(sb.getGreeting())
// }
