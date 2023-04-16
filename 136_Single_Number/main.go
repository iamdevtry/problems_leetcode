package main

func singleNumber(nums []int) int {
	var result int
	for _, num := range nums {
		result ^= num
	}
	return result
}

func main() {
	nums := []int{4, 1, 2, 1, 2}
	println(singleNumber(nums))
}
