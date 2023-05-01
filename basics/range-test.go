package main

import "fmt"

func rangeTeste() {
	slice := []string{"Rafael", "Lino"}
	for idx, value := range slice {
		fmt.Println("idx:", idx, "value:", value)
		for _, innerValue := range value {
			fmt.Printf("value: %q\n", innerValue) // print the rune and not the ascii code
		}
	}
}
