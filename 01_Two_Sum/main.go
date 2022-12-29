package main

import "fmt"

func twoSum(nums []int, target int) []int {
	var rs []int
	for i := 0; i < len(nums); {
		for j := i + 1; j < len(nums); {
			if nums[i]+nums[j] == target {
				rs = append(rs, i)
				rs = append(rs, j)
				break
			}
			j++
		}
		if len(rs) > 0 {
			break
		}
		i++
	}
	return rs
}

func main() {
	rs := twoSum([]int{2, 5, 5, 11}, 10)
	fmt.Println(rs)
}
