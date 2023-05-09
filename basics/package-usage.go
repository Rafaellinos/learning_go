package basics

import (
	"fmt"
	brutils "github.com/Rafaellinos/br-utils"
)

func ExamplePackage() {
	teste := brutils.FormatCpfCnpj("42223485898")
	fmt.Println(teste)
}
