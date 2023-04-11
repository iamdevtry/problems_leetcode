package main

import "fmt"

func maxProfit(price []int) int {
	profit := 0
	n := len(price)
	min, max := 0, price[n-1]

	for i := 0; i < n; i++ {
		min = price[i]
		for j := n - 1; j > 0; j-- {
			max = price[j]
			if max-min > profit && j > i {
				max = price[j]
				profit = max - min
			}
		}
	}

	return profit
}

func main() {
	arr := []int{7, 1, 5, 3, 6, 4}
	// arr := []int{7, 6, 4, 3, 1}
	// arr := []int{4, 1, 5, 2, 7}
	// arr := []int{2, 4, 1}
	fmt.Println(maxProfit(arr))
}
