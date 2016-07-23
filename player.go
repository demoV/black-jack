package blackjack
import "fmt"
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
		fmt.Println(firstValueOfCard, seconValueOfCard, aCard)
		if(minValueOfA){
			sum += firstValueOfCard
		}else{
			sum += seconValueOfCard
		}
	}
	fmt.Println(p.handCards)
	return sum
}