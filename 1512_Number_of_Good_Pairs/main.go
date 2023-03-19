package main

func numIdenticalPairs(nums []int) int {
	result := 0
	existNums := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		count, isExist := existNums[nums[i]]
		if isExist {
			result += count
			existNums[nums[i]] = count + 1
		} else {
			existNums[nums[i]] = 1
		}
	}

	return result
}
