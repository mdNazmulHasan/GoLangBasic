package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func newDeck() deck {
	cards := deck{}
	cardSuites := []string{"spades", "diamonds", "hearts", "clubs"}
	cardValue := []string{"ace", "two", "three", "four"}
	for _, suites := range cardSuites {
		for _, value := range cardValue {
			cards = append(cards, value+" of "+suites)
		}
	}
	return cards
}
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]

}
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}
func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}
