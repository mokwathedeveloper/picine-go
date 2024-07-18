package main

import "fmt"

func Max(a []int) int {

	if len(a) == 0 {
		return 0
	}
	maxValue := a[0]

	for _, value := range a {
		if value > maxValue {
			maxValue = value
		}
	}
	return maxValue

}

func main() {
	a := []int{23, 123, 1, 11, 55, 93}
	max := Max(a)
	fmt.Println(max)
}
