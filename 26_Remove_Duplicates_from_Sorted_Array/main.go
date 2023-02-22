package main

import "fmt"

//Given an integer arrays `nums` sorted in non-decreasing order
//Remove the duplicates in-place such that each unique element appears only once.
//The relative order of the elements should be kept the same.

//Since it is impossible to change the length of the array in some languages,
//you must instead have the result be placed in the first part of the array `nums`.
//More formally, if there are k elements after removing the duplicates,
//then the first k elements of `nums` should hold the final result.
//It does not matter what you leave beyond the first k elements.

//Return `k` after placing the final result in the first k slots of `nums`.

//Do not allocate extra space for another array.
//You must do this by modifying the input array in-place with O(1) extra memory.

func removeDuplicates(nums []int) (int, []int) {
	//if the length of the array is 0, return 0
	if len(nums) == 0 {
		return 0, nums
	}
	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1, nums
}

// Big O: O(n) time | O(1) space

func main() {
	arr := []int{1, 1, 2, 3, 3, 4, 5}

	fmt.Println(removeDuplicates(arr))
	//Output: 5 [1 2 3 4 5 4 5]
}
