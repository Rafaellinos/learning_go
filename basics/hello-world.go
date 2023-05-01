package main

import (
	"fmt"
)

func main() {
	result := calculateAge(1990, 4, 14)
	fmt.Println("Your age is:", result)
	a, d, _ := returnManyExample() // _ is ignored
	result = a + d

	res := reverseString("Rafael")
	fmt.Println("reversed name is: ", res)

	fmt.Println("decimal result:", binaryToDecimal("101010"))

	goWhileLoop(10)
	goBreakLoops(10)
	diceRoller(2, 6, 3)

	getSampleStructures()

	rec := Rectangle{
		5,
		5,
	}
	fmt.Println("Perimeter of 5x5 Rectangle is:", calculateRectanglePerimeter(rec))
	fmt.Println("Area of 5x5 Rectangle is:", calculateRectangleArea(rec))
	arraysExample()
	exerciseArray()
	// slices
	sliceTesting()
	codingExercise()
	examplePackage()
}
