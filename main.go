package main

import (
	"github.com/k0kubun/pp/v3"
)

func main() {
	cards := newDeck()
	pp.Println(cards)
}
