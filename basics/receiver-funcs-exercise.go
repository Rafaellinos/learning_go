package basics

import "fmt"

type PlayerType struct {
	health, energy, maxHealth, maxEnergy uint64
	name                                 string
}

func (p *PlayerType) setHealth(health uint64) {
	if health > p.maxHealth {
		panic(fmt.Sprintf("Cannot apply more than %d health to player %s", p.maxHealth, p.name))
	}
	p.health = health
}

func (p *PlayerType) setEnergy(energy uint64) {
	if energy > p.maxEnergy {
		panic(fmt.Sprintf("Cannot apply more than %d energy to player %s", p.maxEnergy, p.name))
	}
	p.energy = energy
}

func (p PlayerType) String() string {
	return fmt.Sprintf("Name: %s, health: %d", p.name, p.health)
}

func MainReceiverFuncExercise() {
	p := PlayerType{
		health:    0,
		energy:    0,
		maxHealth: 100,
		maxEnergy: 100,
		name:      "Player one",
	}
	p.setHealth(50)
	p.setEnergy(50)
	fmt.Println(p)
}
