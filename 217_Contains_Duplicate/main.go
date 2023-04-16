package main

import "fmt"

func containsDuplicate(nums []int) bool {
	elements := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		elements[nums[i]]++
		if elements[nums[i]] > 1 {
			return true
		}
	}

	return false
}

func main() {
	// arr := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	arr := []int{1, 2, 3, 1}
	// arr := []int{1, 2, 3, 4}
	fmt.Println(containsDuplicate(arr))
}
