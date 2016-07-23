package blackjack

func standerdPack() []Card {
	cardsName := [13]string{"A","2","3","4","5","6","7","8","9","10","J","Q","K"}
	cards := []Card{}
	for i := 0; i < 13; i++ {
		cardName  := cardsName[i]
		for j := 0; j < 4; j++ {
			cards = append(cards, Card{cardName, i+1})
		}
	}
	return cards
}

type Card struct {
	name string
	rank int
}

func (c Card)Value() (int,int){
	rank := c.rank
	if(rank == 1){
		return 1,11
	}
	if(rank > 1 && rank < 9){
		return rank,rank
	}else if(rank >= 10 || rank == 1){
		return 10,10
	}
	return 0,0
}