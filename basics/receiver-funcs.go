package main

import "fmt"

/*
Receiver func Pointer

Funcs associated with structs
Simular to OOP methods
Can modify a struct

e.g.:

type Coordinate struct {
	X, Y int
}

// receiver func
// data struc that operates / func name and params
func (coord *Coordinate) shiftBy(x, y int) {
	coord.X += x
	coord.Y += y
}

// normal func
func shiftBy(x, y, int, coord *Coordinate) {
	coord.X += x
	coord.Y += y
}

coord := Coordinate{5, 5}
coord.shiftBy(1, 1) // result: 6, 6
shiftBy(1, 1) // result 7, 7
*/

///////////////////

/*
Receiver func Value

- Cannot modify a struct

type Coordinate struct {
	X, Y int
}

// no pointer, returns a new coordinate
func (c Coordinate) ShiftDist(other Coordinate) Coordinate {
	return Coordinate{other.X - c.X, other.Y - c.Y}
}

fist := Coordinate{2,2}
second := Coordinate{1, 5}
distance := first.ShiftDist(second) // {-1, 3}
*/

type Space struct {
	occupied bool
}

type ParkingLot struct {
	spaces []Space
}

func occupySpace(lot *ParkingLot, spaceNum int) {
	lot.spaces[spaceNum-1].occupied = true
}

func (park *ParkingLot) occupySpace(spaceNum int) {
	park.spaces[spaceNum-1].occupied = true
}

func (park *ParkingLot) emptySpace(spaceNum int) {
	park.spaces[spaceNum-1].occupied = false
}

func mainReceiverTests() {
	spaces := make([]Space, 10)
	pl := ParkingLot{
		spaces: spaces,
	}

	pl.occupySpace(2)
	occupySpace(&pl, 3)
	fmt.Println(pl)
	pl.emptySpace(2)
	fmt.Println(pl)
}
