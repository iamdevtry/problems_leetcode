package main

// func numWaterBottles(numBottles int, numExchange int) int {
// 	var result int
// 	var numBottlesEmpty int

// 	result = numBottles
// 	numBottlesEmpty = numBottles

// 	for numBottlesEmpty >= numExchange {
// 		numBottles = numBottlesEmpty / numExchange
// 		numBottlesEmpty = numBottlesEmpty % numExchange

// 		numBottlesEmpty += numBottles
// 		result += numBottles
// 	}

// 	return result
// }

func numWaterBottles(numBottles int, numExchange int) int {
	return numBottles + (numBottles-1)/(numExchange-1)
}

func main() {
	println(numWaterBottles(15, 4))
}
