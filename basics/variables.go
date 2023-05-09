package basics

import "fmt"

func variablesTest() {
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
}
