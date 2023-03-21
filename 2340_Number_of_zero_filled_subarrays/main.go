package main

func zeroFilledSubarray(nums []int) int64 {
	var ans int64
	var count int
	for _, v := range nums {
		if v == 0 {
			count++
		} else {
			ans += int64(count * (count + 1) / 2)
			count = 0
		}
	}
	ans += int64(count * (count + 1) / 2)
	return ans
}

func main() {
	nums := []int{1, 3, 0, 0, 2, 0, 0, 4}
	println(zeroFilledSubarray(nums))
}
