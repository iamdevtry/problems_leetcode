package main

import "fmt"

func kidsWithCandies(candies []int, extraCandies int) []bool {
	result := []bool{}
	max := candies[0]
	for i := 0; i < len(candies); i++ {
		if max < candies[i] {
			max = candies[i]
		}
	}

	for i := 0; i < len(candies); i++ {
		if candies[i]+extraCandies >= max {
			result = append(result, true)
		} else {
			result = append(result, false)
		}
	}

	return result
}

func main() {
	candies := []int{12, 1, 12}

	fmt.Println(kidsWithCandies(candies, 10))
}
