package main

import "fmt"

// slices are just a metadata of an existing array, using metadata (address pointing to existing array)
/*
mySlice := []int{1, 2, 3}
item1 := mySlice[0] // accessing first element

slice or array [a,b] a = inclusive, b = exclusive (b not included)
e.g. slice [0:2] -> result == 0,1
[:] entire array

pretty much like python slicing

part1 := []int{1,2,3} // creating slice
part2 := []int{4,5,6} // another slice

part2 = append(part2, 7, 8) // need to re-assign
// [4,5,6,7,8]
combined := append(part1, part2...) // passing the whole slice
// [1,2,3,4,5,6,7,8]

*/

func sliceTesting() {
	slice := make([]int, 10)
	fmt.Println(slice) // [0 0 0 0 0 0 0 0 0 0]
	// multidimensional slices
	board := [][]string{
		[]string{"_", "_", "_"}, // type is option
		{"_", "_", "_"},
		{"_", "_", "_"},
	}
	fmt.Println(board)
}

type Part string

func showLine(line []Part) {
	for idx, _ := range line {
		fmt.Println("part: ", idx)
	}
}

func codingExercise() {
	assemblyLine := make([]Part, 3)
	showLine(assemblyLine)
	assemblyLine = append(assemblyLine, make([]Part, 2)...)
	showLine(assemblyLine)
	assemblyLine = assemblyLine[3:]
	showLine(assemblyLine)
}
