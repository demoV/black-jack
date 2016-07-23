package blackjack

import ("testing")

func TestStanderdPackGives52Cards(t *testing.T) {
	cards := standerdPack()
	numberOfCards := len(cards)
	if numberOfCards != 52 {
		t.Errorf("Expected 52 cards in a pack, but got %v", numberOfCards)
	}
}

func TestFirstCardOfStanderdPackShouldAce(t *testing.T) {
	cards := standerdPack()
	aCard := cards[0]
	valueOfCard,_ := aCard.Value()
	if valueOfCard != 1 {
		t.Errorf("Expected 1, but got %v", valueOfCard)
	}
}

func TestCardCanGiveHisValue(t *testing.T){
	cardJack := Card{"J", 11}
	valueJack1,valueJack2 := cardJack.Value()	
	if (valueJack1 != 10) {	
		t.Errorf("Expected 10 value of jack, but got %v", valueJack1)	
	}
	if (valueJack2 != 10) {	
		t.Errorf("Expected 10 value of jack, but got %v", valueJack1)	
	}

	cardKing := Card{"K", 13}
	value1OfKing,value2OfKing := cardKing.Value()
	if (value1OfKing != 10) {	
		t.Errorf("Expected 10 value of king, but got %v", value1OfKing)	
	}
	if (value2OfKing != 10) {	
		t.Errorf("Expected 10 value of jack, but got %v", value2OfKing)	
	}
}

func TestCardAceHasTwoValue1And11(t *testing.T){
	cardAce := Card{"Ace", 1}
	valueAce1,valueAce2 := cardAce.Value()	
	if (valueAce1 != 1) {	
		t.Errorf("Expected 1 value of jack, but got %v", valueAce1)	
	}
	if (valueAce2 != 11) {	
		t.Errorf("Expected 11 value of jack, but got %v", valueAce1)	
	}
}