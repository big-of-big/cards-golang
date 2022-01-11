package main

func main() {
	cards := newDeck()
	hand, remainingCards := deal(cards, 5)

	hand.print()
	remainingCards.print()

	//slice := []string{"a", "b", "c", "d", "e", "f", "g"}

	// 頭から2つ
	//fmt.Println(slice[0:2]) // [a b]
	//fmt.Println(slice[:2])  // [a b]
	// インデックス番号が2以上全て
	//fmt.Println(slice[2:]) // [c d e f g]
}
