package main

func countOdds(low int, high int) int {
	result := 0
	elements := high - low + 1

	result = elements / 2

	if elements%2 == 1 && high%2 == 1 {
		result++
	}
	return result
}
