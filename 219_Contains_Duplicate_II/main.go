package main

import "fmt"

func containsNearbyDuplicate(nums []int, k int) bool {
	numsTemp := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		j, isExist := numsTemp[nums[i]]

		if isExist && i-j <= k {
			return true
		}

		numsTemp[nums[i]] = i
	}

	return false
}

func main() {
	arr := []int{1, 2, 3, 1}

	fmt.Println(containsNearbyDuplicate(arr, 3))
}
