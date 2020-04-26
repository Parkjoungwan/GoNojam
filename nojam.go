package main

import (
	"fmt"
	"math"
)

var ans int

func nojam9663() {
	var N int
	fmt.Scan(&N)
	for i := 1; i <= N; i++ {
		var Col []int = make([]int, 15)
		Col[1] = i
		nojam9633dfs(1, N, Col)
	}
	fmt.Println(ans)
}

func nojam9633dfs(Row int, N int, Col []int) {
	if Row == N {
		ans++
	} else {
		for i := 1; i <= N; i++ {
			Col[Row+1] = i
			if nojam9633Possible(Row+1, Col) {
				nojam9633dfs(Row+1, N, Col)
			} else {
				Col[Row+1] = 0
			}
		}
	}
}

func nojam9633Possible(C int, Col []int) bool {
	for i := 1; i < C; i++ {
		if Col[i] == Col[C] {
			return false
		}
		if math.Abs(float64(Col[i]-Col[C])) == math.Abs(float64(i-C)) {
			return false
		}
	}
	return true
}

func main() {
	nojam9663()
}
