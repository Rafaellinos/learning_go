package basics

import "fmt"

func sum(nums ...int) int {
	var sum = 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func MainVariadics() {
	fmt.Println(sum(1, 2, 3, 45))
	a := []int{233, 112, 1}
	fmt.Println(sum(a...))
}
