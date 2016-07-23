package blackjack

type Player struct {
	name string
	handCards []Card
}

func (p *Player)Take(card Card){
	p.handCards = append(p.handCards, card)
}

func (p *Player)Sum(minValueOfA bool) int {
	sum := 0
	for _,aCard := range p.handCards{
		firstValueOfCard,seconValueOfCard := aCard.Value()
		if(minValueOfA){
			sum += firstValueOfCard
		}else{
			sum += seconValueOfCard
		}
	}
	return sum
}