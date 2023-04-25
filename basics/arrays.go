package main

import "fmt"

/*
Arrays are fixed-size
same type
may have default values
*/

type House struct {
	windows int
	doors   int
	name    string
}

func arraysExample() {

	//var myArray [3]int // number of items [3]
	myArray := [3]int{1, 2, 3}
	fmt.Println(myArray)
	myArray2 := [...]int{3, 2, 1}
	// 3 dots look at the size on the right size and then, applies
	fmt.Println(myArray2)
	fmt.Println(myArray2[2])
	myArray2[2] = 33
	item2 := myArray[2]
	fmt.Println(item2)
	fmt.Println(len(myArray))

	richHouses := [...]House{
		{30, 30, "white house"},
		{1, 1, "poor house"},
		{0, 0, "funny house"},
	}

	for _, house := range richHouses {
		if house.windows == 0 && house.doors == 0 {
			fmt.Println("No way out from", house.name)
		} else {
			fmt.Println("U can exit the", house.name)
		}
	}

}
