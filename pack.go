package blackjack

type Pack struct {
	packCards []Card
}

func (pack *Pack)Top() Card{
	card := pack.packCards[0]
	pack.packCards = pack.packCards[1:]
	return card
}

func (pack *Pack)Length() int{
	return len(pack.packCards)
}