package blackjack

import "testing"

func TestPlayerCanPickACardAndCanGiveSumOfThoseCardsValue(t *testing.T) {
	player := Player{"Name", make([]Card,0)}
	player.Take(Card{"A",1})
	sum := player.Sum(true)
	if(sum != 1) {
		t.Errorf("Expected to sum is 1, but got %v", sum)
	}
}

func TestPlayerCanGiveSumOfCardsValue(t *testing.T) {
	player := Player{"Name", make([]Card,0)}
	player.Take(Card{"A",1})
	player.Take(Card{"9",9})
	player.Take(Card{"J",11})

	sum := player.Sum(true)
	if(sum != 20) {
		t.Errorf("Expected to sum is 21, but got %v", sum)
	}
}