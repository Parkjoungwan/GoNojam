package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
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
	var s string
	var sum int64
	sc := bufio.NewScanner(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)
	digit := make([]int, 10)
	const maxCapacity = 512 * 1024
	buf := make([]byte, maxCapacity)
	sc.Buffer(buf, maxCapacity)
	defer wr.Flush()
	if sc.Scan() {
		s = sc.Text()
	}
	for i := 0; i < len(s); i++ {
		var tmp = int64(s[i] - '0')
		digit[tmp]++
		sum += tmp
	}
	if digit[0] == 0 || sum%3 != 0 {
		fmt.Fprintf(wr, "-1")
	} else {
		for i := 9; i >= 0; i-- {
			for j := 0; j < digit[i]; j++ {
				fmt.Fprintf(wr, "%d", i)
			}
		}
	}
}
func nojam1929() {
	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()
	var N, M int
	fmt.Scan(&N, &M)
	var PN [1000001]int
	PN[1] = 1
	for i := 2; i <= M; i++ {
		if PN[i] == 0 {
			for j := i * i; j <= M; j += i {
				PN[j] = 1
			}
		}
	}
	for i := N; i <= M; i++ {
		if PN[i] != 1 {
			fmt.Fprintln(wr, i) //이거 출력이 안될지도?
		}
	}
}

func nojam9020() {
	wr := bufio.NewWriter(os.Stdout)
	sc := bufio.NewScanner(os.Stdin)
	defer wr.Flush()
	var T int
	fmt.Scan(&T)
	var PN [10001]int
	var Gold int
	var gold1, gold2 int
	for i := 0; i < T; i++ {
		check := 10001.0
		if sc.Scan() {
			Gold, _ = strconv.Atoi(sc.Text())
		}
		for i := 2; i <= Gold; i++ {
			if PN[i] == 0 {
				for j := i * i; j <= Gold; j += i {
					PN[j] = 1
				}
			}
		}
		for j := 2; j < Gold; j++ {
			for k := 2; k < Gold; k++ {
				if PN[j] != 1 && PN[k] != 1 && j+k == Gold {
					if check > math.Abs(float64(k-j)) {
						gold1 = j
						gold2 = k
						check = math.Abs(float64(k - j))
					}
				}
			}
		}
		if gold1 > gold2 {
			gold1, gold2 = gold2, gold1
		}
		fmt.Fprintf(wr, "%d %d\n", gold1, gold2)
		check = 10001
	}
}

func main() {
	nojam9020()
}
