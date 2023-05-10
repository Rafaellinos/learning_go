package intermediate

import "fmt"

/*
specify the behavior
implicitly implementation
to be implemented, the type must have ALL the interface funcs implemented
Function MUST BE type (POINTER OR VALUE) to be implemented, cannot have different receivers (for consistency)
Name convention for Interfaces it have a suffix "er" at the end. E.g.: Preparer, Resetter, DbHandler
*/

/*
// School has implemented Interface
func FuncInterfaceParameterToType(i Interface) {
	if school, ok := i.(School); ok {
		doSomethingWithSchool(school)
	} else {
		// 'i' has implemented Interface, but it isn't of type School
	}
}
*/

type DemoInterface interface {
	Function1()
	Function2(x int) int
}

type DemoType1 int

func (dt DemoType1) Function1() {
	fmt.Println("Someone called func1 from Demointerface")
}

func (dt DemoType1) Function2(x int) int {
	return x * 2
}

func executeInterfaceDemo(i DemoInterface) {
	// interface parameter can receive a pointer or copy.
	// If use *DemoInterface, will only accept pointers.
	i.Function1()
}

func Function1(i DemoInterface) {
	i.Function1()
}

func MainInterface() {
	DemoType1(14).Function1()
	DemoType1(15).Function2(2)
	executeInterfaceDemo(DemoType1(2))
	aPointer := DemoType1(3)
	executeInterfaceDemo(&aPointer)
	Function1(aPointer)
}
