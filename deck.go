package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

//create a new type of deck whick is slice of string

type deck []string

//'print' is a function with a receiver of type 'deck' and prints all the cards with index
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func newDeck() deck {
	cards := deck{}
	cradSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cradSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

//'deal' is a function  that requires arguments of type 'deck' and 'int', then returns two values of type (deck, deck)
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	//converting deck > slice > sting
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 777)
}

func readDeckFromFile(filePath string) deck {
	bs, err := ioutil.ReadFile(filePath)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	cardSlice := strings.Split(string(bs), ",")

	return deck(cardSlice)
}

func (d deck) suffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
