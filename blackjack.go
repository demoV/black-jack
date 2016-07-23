package blackjack

func standerdPack() [52]Card {
	cards := [52]Card{}
	for i := 2; i < 15; i++ {
		for j := 0; j < 4; j++ {
			
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