package tests

import (
	"learning_go/basics"
	"testing"
)

func TestPlayerSetAboveMax(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Function did not panicked.")
		}
	}()

	p := basics.PlayerType{
		Health:    0,
		Energy:    0,
		MaxHealth: 100,
		MaxEnergy: 100,
		Name:      "Player one",
	}

	p.SetEnergy(5000)
}

func TestPlayerSetSuccess(t *testing.T) {
	playerExpectHealth := uint64(90)
	p := basics.PlayerType{
		Health:    0,
		Energy:    0,
		MaxHealth: 100,
		MaxEnergy: 100,
		Name:      "Player one",
	}

	p.SetHealth(playerExpectHealth)

	if p.Health != playerExpectHealth {
		t.Errorf("Player health %v != %v", p.Health, playerExpectHealth)
	}

}
