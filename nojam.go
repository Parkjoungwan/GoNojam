package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func nojam11650() {
	sc := bufio.NewScanner(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	var N int
	if sc.Scan() {
		N = toInt(sc.Bytes())
	}
	arr := make([][2]int, N)
	for i := 0; i < N; i++ {
		if sc.Scan() {
			// tmp, tmpA, tmpB는 값을 받기 위한 임시 변수
			tmp := strings.Fields(sc.Text())
			tmpA, _ := strconv.Atoi(tmp[0])
			tmpB, _ := strconv.Atoi(tmp[1])
			arr[i] = [2]int{tmpA, tmpB}
		}
	}
	sort.Sort(cusheap(arr))
	// fmt.Fprintf(wr, "\n")
	for i := 0; i < N; i++ {

		fmt.Fprintf(wr, "%d %d\n", arr[i][0], arr[i][1])
	}
}

func toInt(buf []byte) (n int) {
	for _, v := range buf {
		n = n*10 + int(v-'0')
	}
	return
}

type cusheap [][2]int

func (h cusheap) Len() int { return len(h) }
func (h cusheap) Less(i, j int) bool {
	if h[i][0] == h[j][0] {
		return h[i][1] < h[j][1]
	} else {
		return h[i][0] < h[j][0]
	}
}

func (h cusheap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func nojam16212() {
	var N int
	fmt.Scan(&N)
	X := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&X[i])
	}
	sort.Sort(sort.IntSlice(X))
	for i := 0; i < N; i++ {
		fmt.Print(X[i], " ")
	}
}

func nojam1059() {
	var N int
	var num [1001]int
	var Luckey int
	fmt.Scan(&N)
	for i := 0; i < N; i++ {
		fmt.Scan(&Luckey)
		num[Luckey] = 1
	}
	var Unlukey int
	fmt.Scan(&Unlukey)
	var down int
	down = 0
	var up int
	var check int
	for i := 1; i < 1001; i++ {
		if num[i] == 1 && check == 0 {
			down = i
		}
		if i == Unlukey {
			check = 1
		}
		if num[i] == 1 && i == Unlukey {
			fmt.Println(0)
			return
		}
		if check == 1 && num[i] == 1 {
			up = i
			break
		}
	}
	fmt.Println(down, Unlukey, up)
	fmt.Println((Unlukey-down)*(up-Unlukey) - 1)
}
func nojam1041() {
	var N int
	fmt.Scan(&N)
	var A, B, C, D, E, F int
	fmt.Scan(&A, &B, &C, &D, &E, &F)
	Num := make([]int, 6)
	Num[0] = A
	Num[1] = B
	Num[2] = C
	Num[3] = D
	Num[4] = E
	Num[5] = F
	sort.Sort(sort.IntSlice(Num))
	for i := 1; i < 5; i++ {
		Num[0] += Num[i]
	}
	if A > F {
		A = F
	}
	if B > E {
		B = E
	}
	if C > D {
		C = D
	}
	var for3, for2, for1 int
	for3 = A + B + C
	if A >= B && A >= C {
		for2 = B + C
		if B > C {
			for1 = C
		} else {
			for1 = B
		}
	} else if B >= C && B >= A {
		for2 = A + C
		if A > C {
			for1 = C
		} else {
			for1 = A
		}
	} else {
		for2 = A + B
		if A > B {
			for1 = B
		} else {
			for1 = A
		}
	}
	for3 = 4 * for3
	for2 = (8*N - 12) * for2
	for1 = (5*N*N - 16*N + 12) * for1
	ans := for3 + for2 + for1
	if N == 1 {
		fmt.Println(Num[0])
	} else {
		fmt.Println(ans)
	}
}
func main() {
	nojam1041()
}
