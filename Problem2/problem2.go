package main

import (
	"fmt"
)

func ShiftArray(array []int, i int) []int {
	var result []int

	current := [][]int{{array[0], array[1], array[2]}, {array[3], array[4], array[5]}, {array[6], array[7], array[8]}}
	target := [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}

	for j := 0; j < i; j++ {
		target[0][1], target[0][0], target[1][0], target[2][0], target[2][1], target[2][2], target[1][2], target[0][2], target[1][1] =
			current[0][0], current[1][0], current[2][0], current[2][1], current[2][2], current[1][2], current[0][2], current[0][1], current[1][1]

		current, target = target, [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}
	}
	result = []int{current[0][0], current[0][1], current[0][2], current[1][0], current[1][1], current[1][2], current[2][0], current[2][1], current[2][2]}

	return result
}

func main() {
	fmt.Println(ShiftArray([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 1)) // Output: [4 1 2 7 5 3 8 9 6]
	fmt.Println(ShiftArray([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 2)) // Output: [7 4 1 8 5 2 9 6 3]
}
