package main

import "math"

func minReorder(n int, connections [][]int) int {
	graph := make([][]int, n)

	for _, pair := range connections {
		i := pair[0]
		j := pair[1]
		graph[i] = append(graph[i], j)
		graph[j] = append(graph[j], -i)
	}

	return traverse(&graph, 0, -1)
}

func traverse(graph *[][]int, i int, parent int) int {
	changes := 0

	for _, to := range (*graph)[i] {
		target := int(math.Abs(float64(to)))

		if target == parent {
			continue
		}

		if to > 0 {
			changes++
		}

		changes += traverse(graph, target, i)
	}

	return changes
}
