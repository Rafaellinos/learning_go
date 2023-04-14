package main

import (
	"fmt"
)

// rune = alias for int32 that represents a letter
// string is a series bytes | String are not null terminated
// "" <- string
// ” <- rune
// `let's code\n` <- raw string
func main() {
	//fmt.Printf("vai tomar no cu")
	//var example int = 3 //
	//example = 4
	c := "sample"
	var b = "sex"
	c = b
	var name string // default value of "", number default 0 and other nil
	var sex rune    // default 0
	sex = 'a'       // prints the number representation
	fmt.Print(c)
	fmt.Print(name)
	fmt.Print("-----------")
	fmt.Print(sex)
	fmt.Print("-----------")
	j, b := "13", "14" // comma, ok idiom allows to reassign
	fmt.Printf(b)
	fmt.Printf(j)
	// camelCase
	const MaxSpeed = 30 // cannot reassign
	result := calculateAge(1990, 4, 14)
	fmt.Println("Your age is:", result)
	a, d, _ := returnManyExample() // _ is ignored
	result = a + d

	res := reverseString("Rafael")
	fmt.Println("reversed name is: ", res)

	fmt.Println("decimal result:", binaryToDecimal("101010"))
}

// go commands:
// build -> distribute
// build -race -> check for concurrency problems
// run -> runs directly, no output executable
// mod -> manage modules and dependencies , mod tidy -> update dependencies
// test -> runs the projects test
// fmt -> format all source files

//------------

// variables -> Name, data, etc
