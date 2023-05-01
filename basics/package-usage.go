package main

import (
	"fmt"
	brutils "github.com/Rafaellinos/br-utils"
)

func examplePackage() {
	teste := brutils.FormatCpfCnpj("42223485898")
	fmt.Println(teste)
}
