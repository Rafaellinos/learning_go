package basics

import "fmt"

type PlayerType struct {
	Health, Energy, MaxHealth, MaxEnergy uint64
	Name                                 string
}

func (p *PlayerType) SetHealth(health uint64) {
	if health > p.MaxHealth {
		panic(fmt.Sprintf("Cannot apply more than %d health to player %s", p.MaxHealth, p.Name))
	}
	p.Health = health
}

func (p *PlayerType) SetEnergy(energy uint64) {
	if energy > p.MaxEnergy {
		panic(fmt.Sprintf("Cannot apply more than %d energy to player %s", p.MaxEnergy, p.Name))
	}
	p.Energy = energy
}

func (p PlayerType) String() string {
	return fmt.Sprintf("Name: %s, health: %d", p.Name, p.Health)
}

func MainReceiverFuncExercise() {
	p := PlayerType{
		Health:    0,
		Energy:    0,
		MaxHealth: 100,
		MaxEnergy: 100,
		Name:      "Player one",
	}
	p.SetHealth(50)
	p.SetEnergy(50)
	fmt.Println(p)
}
