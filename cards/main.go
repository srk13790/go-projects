package main

// import "fmt"

func main() {
	// var card string = "Ace Of Spades"
	// card := "Ace Of Spades"
	// card := newCard()
	// fmt.Println(card)

	// cards := [] string {newCard(), "Ace of Spades"}
	// cards := newDeck()

	// for _, card := range cards {
	// 	fmt.Println(card)
	// }

	// cards.print()

	// cards.printAnother()

	// fmt.Println(cards)

	// first5, last5 := deal(cards, 5);

	// first5.print()
	// last5.print()

	// greeting := "Hi there!"

	// fmt.Println([]byte(greeting))

	// cards := newDeck()

	// fmt.Println(cards.toString())

	// cards.saveToFile("myCards")


	// cards := newDeckFromFile("myCards")
	// cards.print()

	cards := newDeck()
	cards.shuffle()
	cards.print()

}

func newCard() string {
	return "Five Of Diamonds"
}
