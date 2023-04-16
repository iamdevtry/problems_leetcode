package main

import "fmt"

func majorityElement(nums []int) int {
	values := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		values[nums[i]]++
		if values[nums[i]] > len(nums)/2 {
			return nums[i]
		}

	}

	return -1
}

func main() {
	arr := []int{2, 2, 1, 1, 1, 2, 2}
	fmt.Println(majorityElement(arr))
}
