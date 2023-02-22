package main

// Given an integer `numRows`, return the first numRows of Pascal's triangle.
// In Pascal's triangle, each number is the sum of the two numbers directly
//above it as shown:
//					1
//				1		1
//			1		2		1
//		1		3		3		1
//	1		4		6		4		1

func generate(numRows int) [][]int {
	var result [][]int
	for i := 0; i < numRows; i++ {
		result = append(result, make([]int, i+1))
		result[i][0], result[i][i] = 1, 1
		for j := 1; j < i; j++ {
			result[i][j] = result[i-1][j-1] + result[i-1][j]
		}
	}
	return result
}

func main() {
	generate(5)
}
