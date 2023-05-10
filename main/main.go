package main

import (
	"fmt"
	"learning_go/basics"
)

func init() {
	// each package has its own init and main
	fmt.Println("executed before main")
	/*
		Ex: compile regex before initializing the package and will be ready
		var EmailExpr *regex.Regexp
		func init() {
			compiled, ok := regexp.Compile(`[0-9]`)
			if ok != nil {
				panic("failed to compile")
			}
			EmailExpr = compiled
		}

	*/
}

func main() {
	result := basics.CalculateAge(1990, 4, 14)
	fmt.Println("Your age is:", result)
	a, d, _ := basics.ReturnManyExample() // _ is ignored
	result = a + d

	res := basics.ReverseString("Rafael")
	fmt.Println("reversed name is: ", res)

	fmt.Println("decimal result:", basics.BinaryToDecimal("101010"))

	basics.GoWhileLoop(10)
	basics.GoBreakLoops(10)
	basics.DiceRoller(2, 6, 3)

	basics.GetSampleStructures()

	//rec := basics.Rectangle{
	//	5,
	//	5,
	//}
	//fmt.Println("Perimeter of 5x5 Rectangle is:", basics.CalculateRectanglePerimeter(rec))
	//fmt.Println("Area of 5x5 Rectangle is:", basics.CalculateRectangleArea(rec))
	basics.ArraysExample()
	basics.ExerciseArray()
	// slices
	basics.SliceTesting()
	basics.CodingExercise()
	basics.ExamplePackage()
	basics.RangeTeste()
	basics.MapsExample()
	basics.ExerciseMaps()
	basics.PointersExample()
	basics.LibraryExercise()
	basics.MainReceiverTests()
	basics.MainReceiverFuncExercise()
	basics.MainIotaTest()
	basics.MainVariadics()
}
