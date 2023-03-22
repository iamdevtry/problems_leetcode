package main

import "math"

func find(x int, connectNode []int) int {
	if connectNode[x] != x {
		connectNode[x] = find(connectNode[x], connectNode)
	}
	return connectNode[x]
}

func minScore(n int, roads [][]int) int {
	connectNode, minDistance := make([]int, n+1), make([]int, n+1)

	for i := 1; i <= n; i++ {
		connectNode[i] = i
		minDistance[i] = math.MaxInt16
	}
	for i := 0; i < len(roads); i++ {
		root_a := find(roads[i][0], connectNode)
		root_b := find(roads[i][1], connectNode)
		minPath := math.Min(float64(minDistance[root_a]), float64(minDistance[root_b]))
		minPath = math.Min(minPath, float64(roads[i][2]))
		if root_a > root_b {
			connectNode[root_a] = connectNode[root_b]
		} else {
			connectNode[root_b] = connectNode[root_a]
		}
		minDistance[root_a] = int(minPath)
		minDistance[root_b] = int(minPath)

	}
	return minDistance[find(n, connectNode)]
}

func main() {
}
