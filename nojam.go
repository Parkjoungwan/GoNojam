package main

import (
	"fmt"
	"sort"
)

func nojam2309() {
	Height := make([]int, 9)
	var all int
	var i, j int
	var fake1, fake2 int
	for i := 0; i < 9; i++ {
		fmt.Scan(&Height[i])
	}
	sort.Sort(sort.IntSlice(Height))
	for i = 0; i < 9; i++ {
		all += Height[i]
	}
	for i = 0; i < 9; i++ {
		for j = 0; j < 9; j++ {
			if all-Height[i]-Height[j] == 100 {
				fake1 = i
				fake2 = j
			}
		}
	}
	for k := 0; k < 9; k++ {
		if fake1 == k || fake2 == k {
		} else {
			fmt.Println(Height[k])
		}
	}
}

func main() {
	nojam2309()
}
