package main

import "fmt"

type deck []string


func newDeck() deck{
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds","Hearts","Clubs"}
	cradValues :=[]string{"Ace","Two","Three","Four"}

	for _,s :=range cardSuits{
		for _,v :=range cradValues{
			cards = append(cards, v+" of "+s)
		}
	}
	return cards
}

func (d deck) print(){
	for i, card := range d{
		fmt.Println(i,card)
	}
}

func deal (d deck, handsize int) (deck, deck){

	return d[:handsize], d[handsize:]
	
}