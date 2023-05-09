package basics

import "fmt"

/*
Copy of each function argument is made every time u call a func, regardless the size.
Pointers can change this behavior
Can be slow for large data structures
The value of the variable itself is the memory address
To access the data, use dereferencing
Can change the value because it's using the same memory address

(*) for pointers
(&) create a point based on variable

e.g.:

intValue1 := 10
var intValue *int // indication that this type is pointer
intValue2 = &value // creating a pointer

&value = pointer to value
*int = pointer to integer
*/

type Player struct {
	nickName   string
	hp         int
	strength   int
	liveStatus string
}

func combatPlayers(playerOffender *Player, playerTarget *Player) {
	playerTarget.hp = -playerOffender.strength
	if playerTarget.hp <= 0 {
		playerTarget.hp = 0
		playerTarget.liveStatus = "dead"
	}
}

func PointersExample() {
	//value := 22
	//var valuePointer *int
	//fmt.Println(value)
	//fmt.Println(valuePointer)
	//valuePointer = &value
	//fmt.Println(value)
	nero := Player{
		"Nero",
		100,
		32,
		"alive",
	}
	vergil := Player{
		"Vergil",
		70,
		15,
		"alive",
	}
	combatPlayers(&nero, &vergil)
	fmt.Println(nero)
	fmt.Println(vergil)
}
