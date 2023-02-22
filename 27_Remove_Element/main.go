package main

//Given an interger array nums and an integer val
//remove all occurrences of val in nums in-place.
//The relative order of the elements may be changed.

//Since it is impossible to change the length of the array in some languages,
//you must instead have the result be placed in the first part of the array nums.
//More formally, if there are k elements after removing the duplicates, then the
//first k elements of nums should hold the final result. It does not matter what
//you leave beyond the first k elements.

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}

	i := 0
	for j := 0; j < len(nums); j++ {
		if nums[j] != val {
			nums[i] = nums[j]
			i++
		}
	}
	return i
}

func main() {
	arr := []int{1, 1, 2, 3, 3, 4, 5}
	removeElement(arr, 3)
}
