package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

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

func (d deck) toString() string{
	return strings.Join([]string(d),",")
}

func (d deck) saveToFile (filename string) error{
	 return ioutil.WriteFile(filename, []byte(d.toString()),0666) //0666 means any one can ready and write a file
}

func newDeckFromFile (filename string) deck {
	bs, err := ioutil.ReadFile(filename)

	if err != nil{
		    //Option #1 - Log the error and return a call to newDeck()
			// Option #2 -  Log the error and entirely quit the program
			 fmt.Println("Error: ", err)
			 os.Exit(1)
	}
	s:= strings.Split(string(bs),",")
	return  deck(s)

}
 