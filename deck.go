package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

// スペードからクローバーまでの1〜13の組み合わせの新しいデッキを作る
func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"スペード", "ダイア", "ハート", "クローバー"}
	cardValues := []string{"エース", "2", "3", "4"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, suit+"の"+value)
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// 52枚のカードからなるデッキを5枚と47枚のように分ける
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
	//slice := []string{"a", "b", "c", "d", "e", "f", "g"}

	// 頭から2つ
	//fmt.Println(slice[0:2]) // [a b]
	//fmt.Println(slice[:2])  // [a b]
	// インデックス番号が2以上全て
	//fmt.Println(slice[2:]) // [c d e f g]
}

// deckを[]stringに型変換をして、要素をカンマ区切りで1つ文字列にする ["hoge", "piyo"] -> "hoge,piyo"
func (d deck) toString() string {
	// deckと[]stringは互換性がある。[]string(deck)でdeckから[]stringに型変換できる。
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1) // 0が成功。それ以外はエラー。
	}
	// []byte→string→[]string
	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
