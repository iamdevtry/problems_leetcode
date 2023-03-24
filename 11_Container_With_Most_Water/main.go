package main

func maxArea(height []int) int {
	var result int
	n := len(height)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			currentArea := min(height[i], height[j]) * (j - i)
			result = max(currentArea, result)
		}
	}

	return result
}

// Time Complexity: O(n^2)
// Space Complexity: O(1)

func maxAreaOptimized(height []int) int {
	var result int
	i, j := 0, len(height)-1
	for i < j {
		currentArea := min(height[i], height[j]) * (j - i)
		result = max(result, currentArea)

		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	println(maxArea(height))
	println(maxAreaOptimized(height))
}
