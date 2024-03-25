package main

import "fmt"

func maxProfit(prices []int) int {
	minPrice := prices[0]
	maxP := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if p := prices[i] - minPrice; p > maxP {
			maxP = p
		}
	}
	return maxP
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))
}
