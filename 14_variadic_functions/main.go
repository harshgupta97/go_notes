package main

import "fmt"

func sums(nums ...int) {
	fmt.Println(nums, " ")
	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	sums(1, 2, 3)
	sums(1, 2, 3, 4)

	nums := []int{1, 2, 3, 4}
	sums(nums...)
}
