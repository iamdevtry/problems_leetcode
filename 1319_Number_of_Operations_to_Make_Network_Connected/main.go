package main

import "fmt"

type UnionFind struct {
	parent []int
	count  int
}

func NewUnionFind(n int) *UnionFind {
	uf := &UnionFind{
		parent: make([]int, n),
		count:  n,
	}
	for i := 0; i < n; i++ {
		uf.parent[i] = i
	}
	return uf
}

func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf *UnionFind) Union(x, y int) {
	rootX, rootY := uf.Find(x), uf.Find(y)
	if rootX != rootY {
		uf.parent[rootX] = rootY
		uf.count--
	}
}

func (uf *UnionFind) Count() int {
	return uf.count
}

func makeConnected(n int, connections [][]int) int {
	if len(connections) < n-1 {
		return -1
	}
	uf := NewUnionFind(n)
	for _, c := range connections {
		uf.Union(c[0], c[1])
	}
	return uf.Count() - 1
}

func main() {
	fmt.Println(makeConnected(4, [][]int{{0, 1}, {0, 2}, {1, 2}}))
	fmt.Println(makeConnected(6, [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 2}, {1, 3}}))
	fmt.Println(makeConnected(6, [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 2}}))
	fmt.Println(makeConnected(5, [][]int{{0, 1}, {0, 2}, {3, 4}, {2, 3}}))
}
