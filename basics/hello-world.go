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
}
