package intermediate

import "fmt"

type Preparer interface {
	PrepareDish()
}

type Chicken string

type Salad string

func (c Chicken) PrepareDish() {
	fmt.Println("Preparing chicken...")
}

func (s Salad) PrepareDish() {
	fmt.Println("Preparing salad...")
}

func prepareDishes(dishes []Preparer) {
	fmt.Println("initiating preparation....")
	for _, dishe := range dishes {
		dishe.PrepareDish()
	}
	fmt.Println("Dishes ready to serve")
}

func MainInterfacesDemo() {
	dishes := []Preparer{
		Chicken("C"),
		Salad("S"),
	}
	prepareDishes(dishes)
}
