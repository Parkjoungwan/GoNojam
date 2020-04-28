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

func nojam1699() {
	var number [100001]int
	var N int
	fmt.Scan(&N)
	for i := 1; i <= N; i++ {
		for j := 1; j*j <= i; j++ {
			if number[i] > number[i-j*j]+1 || number[i] == 0 {
				number[i] = number[i-j*j] + 1
			}
		}
	}
	fmt.Println(number[N])
}

func nojam1010() {
	var N, M int
	var T int
	fmt.Scan(&T)
	var DP [31][31]int
	for i := 1; i <= 30; i++ {
		DP[1][i] = i
	}
	for i := 2; i <= 30; i++ {
		for j := i; j <= 30; j++ {
			for k := j; k >= i; k-- {
				DP[i][j] = DP[i][j] + DP[i-1][k-1]
			}
		}
	}
	for i := 0; i < T; i++ {
		fmt.Scan(&N, &M)
		fmt.Println(DP[N][M])
	}
}

func nojam11399() {
	var N int
	fmt.Scan(&N)
	Per := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&Per[i])
	}
	sort.Sort(sort.IntSlice(Per))
	var ans int
	for i := 0; i < N; i++ {
		for j := 0; j <= i; j++ {
			ans += Per[j]
		}
	}
	fmt.Println(ans)
}

func nojam2875() {
	var M int
	var F int
	var K int
	fmt.Scan(&F, &M, &K)
	for K != 0 {
		if F > M*2 {
			F -= 1
			K -= 1
		} else {
			M -= 1
			K -= 1
		}
	}
	var count int
	for M != 0 {
		F -= 2
		if F < 0 {
			break
		}
		M -= 1
		count++
	}
	fmt.Println(count)

}
func nojam10610() {
	var N string
	var checkNum int
	fmt.Scan(&N)
	var checkZero int
	for i := 0; i < len(N); i++ {
		checkNum += int(N[i] - '0')
		if int(N[i]-'0') == 0 {
			checkZero = 1
		}
	}
	if checkNum%3 == 0 && checkZero == 1 {
		fmt.Println(checkNum * 10)
	} else {
		fmt.Println("-1")
	}
}

func main() {
	nojam10610()
}
