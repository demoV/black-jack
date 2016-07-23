package blackjack

import "testing"

func TestPackCanGiveCardFromTop(t *testing.T) {
	cards := standerdPack()
	pack := Pack{cards}
	topCard := pack.Top()
	cardValue,_ := topCard.Value()

	if(cardValue != 1){
		t.Errorf("Expected 1, but got %v", cardValue)
	}
}

func TestLengthOfPackShouldDecreaseByOneAfterPickingTheCard(t *testing.T) {
	cards := standerdPack()
	pack := Pack{cards}
	topCard := pack.Top()
	cardValue,_ := topCard.Value()

	if(cardValue != 1){
		t.Errorf("Expected 1, but got %v", cardValue)
	}
	if(pack.Length() != 51){
		t.Errorf("Pack's length should be 51, but got %v", pack.Length())
	}
}