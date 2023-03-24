package main

import "math"

func reverse(x int) int {
	var result int
	for x != 0 {
		result = result*10 + x%10
		x = x / 10
	}
	if result > math.MaxInt32 || result < -math.MaxInt32 {
		return 0
	}
	return result
}

func main() {
	println(reverse(123))
	println(reverse(-123))
	println(reverse(120))
	println(reverse(1534236469))
}
