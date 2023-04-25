package main

import "fmt"

/* structures simular to class.
data stored in groups
Data point in structure -> field
Functionalities in structures
*/

type Sample struct {
	name         string
	age          int64
	defaultValue string // default value empty string
}

type Rectangle struct {
	width, length int
}

func getSampleStructures() Sample {
	data := Sample{
		name: "Rafael",
		age:  32,
	}
	fmt.Println(data)
	//fmt.Println(data.name)
	// example of anonymous structure
	sample2 := struct {
		fullName string
		age      int64
	}{"Rafael Veloso", 32}
	// must provide all field, use 'var sample struct' to
	// have default values
	fmt.Println(sample2.fullName)
	return data
}

func calculateRectangleArea(rec Rectangle) int {
	return rec.length * rec.width
}

func calculateRectanglePerimeter(rec Rectangle) int {
	return 2 * (rec.length + rec.width)
}
