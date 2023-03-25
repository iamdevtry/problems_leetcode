package main

func binarySearch(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}

func searchInsert(nums []int, target int) int {
	return binarySearch(nums, target)
}

func main() {
	nums := []int{1, 3, 5, 6}
	target := 7
	println(searchInsert(nums, target))
}
