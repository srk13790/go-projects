package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

//crate a new type of 'deck'
// which is a slice of strings

type deck [] string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func newDeck () deck {
	cards := deck{}

	cardSuite := []string {"Spades", "Diamonds", "Hearts", "Clubs"}

	cardValues := []string {"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuite {
		for _, value := range cardValues {
			cards = append(cards, value + " Of " + suit )
		}
		
	}

	return cards
}


func (d1 deck) printAnother() {
	for i, card := range d1 {
		fmt.Println(i, card, "Checking")
	}
}

func deal (d deck, handSize int ) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
	
}

func (d deck) saveToFile (fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

func newDeckFromFile (fileName string) deck {
	bs, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	} 

	s := strings.Split(string(bs), ",")

	return  deck(s);
}

func (d deck) shuffle () {

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}