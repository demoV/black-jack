package blackjack

import "testing"

func TestPlayerCanPickACardAndCanGiveSumOfThoseCardsValue(t *testing.T) {
	player := Player{"Name", make([]Card,0)}
	player.Take(Card{"A",1})
	sum,_ := player.Sum(true)
	if(sum != 1) {
		t.Errorf("Expected to sum is 1, but got %v", sum)
	}
}

func TestPlayerCanGiveSumOfCardsValue(t *testing.T) {
	player := Player{"Name", make([]Card,0)}
	player.Take(Card{"A",1})
	player.Take(Card{"9",9})
	player.Take(Card{"J",11})

	sum,_ := player.Sum(true)
	if(sum != 20) {
		t.Errorf("Expected to sum is 21, but got %v", sum)
	}
}

func TestSumOfCardsIsGreaterThan21ItShouldReturnTrue(t *testing.T) {
	player := Player{"Name", make([]Card,0)}
	player.Take(Card{"A",1})
	player.Take(Card{"9",9})
	player.Take(Card{"K",13})
	player.Take(Card{"2",2})

	sum,bust := player.Sum(true)
	if(!bust){
		t.Errorf("Expected to bust if sum is greater then 21,sum = %v", sum)
	}
	
}