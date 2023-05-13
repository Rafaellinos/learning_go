package intermediate

import "fmt"

type Vehicle interface {
	GetSize() int
}

type Motorcycle struct {
	size int
}

type Truck struct {
	size int
}

type Car struct {
	size int
}

func (m Motorcycle) GetSize() int {
	return m.size
}

func (t Truck) GetSize() int {
	return t.size
}

func (c Car) GetSize() int {
	return c.size
}

type Lifter interface {
	Lift(v Vehicle)
}

type HeavyLifter string

type SmallLifter string

type StandardLifter string

func (hl HeavyLifter) Lift(v Vehicle) {
	fmt.Println("Lifting Heavy")
}

func (sl StandardLifter) Lift(v Vehicle) {
	fmt.Println("Lifting Standard")
}

func (sml SmallLifter) Lift(v Vehicle) {
	fmt.Println("Lifting Small")
}

func doLift(v Vehicle) {
	if v.GetSize() <= 5 {
		SmallLifter("SM").Lift(v)
	}
	if v.GetSize() > 5 && v.GetSize() <= 10 {
		StandardLifter("ST").Lift(v)
	}
	if v.GetSize() > 10 {
		HeavyLifter("HV").Lift(v)
	}
}

func MainInterfacesExercise() {
	car := Car{
		size: 10,
	}
	doLift(car)
	bike := Motorcycle{
		size: 2,
	}
	doLift(bike)
	truck := Truck{
		size: 1000,
	}
	doLift(truck)
}
