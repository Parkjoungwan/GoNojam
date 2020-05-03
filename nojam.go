package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	sc = bufio.NewScanner(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)
)

/*func nojam10799() { //아직 못 품 질문 해 놓았다.
	var Plan string
	var result int
	var stick [100001]int
	j := 0
	tail := 0

	if sc.Scan() {
		Plan = sc.Text()
	}
	Plan = strings.ReplaceAll(Plan, "()", "a")
	for i := 0; i < len(Plan); i++ {
		if Plan[i] == '(' {
			j++
		} else if Plan[i] == 'a' {
			if j != 0 {
				for k := 0; k < j; k++ {
					if stick[k] != -1 {
						stick[k]++
					}
				}
			}
		} else {
			for k := 0; k < j; k++ {
				if stick[k] != -1 {
					tail = k
				}
			}
			result += stick[tail] + 1
			stick[tail] = -1
		}
	}

	fmt.Println(result)
}*/
func delete(i int, s []int) []int {
	s = append(s[:i], s[i+1:]...)
	c := make([]int, len(s))
	copy(c, s)
	return c
}

func nojam11866() {
	var N, K int
	defer wr.Flush()
	fmt.Scan(&N, &K)
	arr := make([]int, N)
	for i := 0; i < N; i++ {
		arr[i] = i + 1
	}
	var val = K - 1
	fmt.Fprintf(wr, "<")
	for i := 0; i < N; i++ {
		fmt.Fprintf(wr, "%d", arr[val])
		if i != N-1 {
			fmt.Fprintf(wr, ", ")
		}
		arr = delete(val, arr)
		val += K - 1
		if N-i-1 != 0 {
			val %= N - i - 1
		}
	}
	fmt.Fprintf(wr, ">")
}
func test() {
	var N int
	fmt.Scan(&N)
	arr := make([]int, N)
	for i := 0; i < N; i++ {
		arr[i] = i + 1
	}
	fmt.Println(arr)
	arr = delete(6, arr)
	fmt.Println(arr)
}
func main() {
	nojam11866()
}
