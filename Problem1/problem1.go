package main

import (
	"fmt"
	"math"
)

func maxProfit(prices []int, i int) int {
	if len(prices) == 0 {
		return 0
	}

	currentTransactions := make([]int, len(prices))
	previousTransactions := make([]int, len(prices))

	for t := 1; t <= i; t++ {
		var current, previous []int
		if t%2 == 0 {
			current = currentTransactions
			previous = previousTransactions
		} else {
			current = previousTransactions
			previous = currentTransactions
		}

		maxPrev := math.MinInt32

		for i := 1; i < len(prices); i++ {
			maxPrev = int(math.Max(float64(maxPrev), float64(previous[i-1]-prices[i-1])))
			current[i] = int(math.Max(float64(current[i-1]), float64(maxPrev+prices[i])))
		}
	}

	if i%2 == 0 {
		return currentTransactions[len(prices)-1]
	} else {
		return previousTransactions[len(prices)-1]
	}
}

func main() {
	prices := []int{4, 11, 2, 20, 59, 80}
	i := 2
	maxProfit := maxProfit(prices, i)
	fmt.Println(maxProfit)
}
