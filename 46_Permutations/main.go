package main

func backtrack(nums []int, path []int, used []bool, result *[][]int) {
	if len(path) == len(nums) {
		*result = append(*result, append([]int{}, path...))
		return
	}

	for i := 0; i < len(nums); i++ {
		if used[i] {
			continue
		}
		path = append(path, nums[i])
		used[i] = true
		backtrack(nums, path, used, result)
		path = path[:len(path)-1]
		used[i] = false
	}
}

func permute(nums []int) [][]int {
	var result [][]int
	var path []int
	used := make([]bool, len(nums))
	backtrack(nums, path, used, &result)
	return result
}
