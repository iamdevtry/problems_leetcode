package main

import (
	"fmt"
)

type rangeNums struct {
	start int
	end   int
}

func summaryRanges(nums []int) []string {
	result := []string{}

	tmp := make(map[int]rangeNums)

	j := 0
	for i := 0; i < len(nums); i++ {
		value, _ := tmp[j]

		if nums[i]-1 == value.end {
			rangeN := rangeNums{value.start, nums[i]}
			tmp[j] = rangeN
		}
		rangeN := rangeNums{nums[i], nums[i]}
		j++
		tmp[j] = rangeN
	}

	fmt.Println(tmp)
	// for _, v := range tmp {
	// 	result = append(result, v)
	// }

	return result
}

func main() {
	arr := []int{0, 2, 3, 4, 6, 8, 9}
	fmt.Println(summaryRanges(arr))
}
