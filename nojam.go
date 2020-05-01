package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var (
	sc    = bufio.NewScanner(os.Stdin)
	wr    = bufio.NewWriter(os.Stdout)
	n     int
	a     [11]int
	op    [4]int
	ops   [10]int
	check [10]bool
	idx   int
	min64 = 1<<63 - 1
	max64 = -1 << 63
)

func nojam9020() {
	defer wr.Flush()
	var T int
	if sc.Scan() {
		T, _ = strconv.Atoi(sc.Text())
	}
	var PN [10001]int
	var Gold int
	var gold1, gold2 int
	for i := 2; i <= 10000; i++ {
		if PN[i] == 0 {
			for j := i * i; j <= 10000; j += i {
				PN[j] = 1
			}
		}
	}
	for i := 0; i < T; i++ {
		check := 10001.0
		if sc.Scan() {
			Gold, _ = strconv.Atoi(sc.Text())
		}
		for j := 2; j < Gold; j++ {
			if PN[j] != 1 && PN[Gold-j] != 1 {
				if check > math.Abs(float64(Gold-2*j)) {
					gold1 = j
					gold2 = Gold - j
					check = math.Abs(float64(Gold - 2*j))
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
func nojam14888() {
	if sc.Scan() {
		n = toInt(sc.Bytes())
	}

	if sc.Scan() {
		tmp := bytes.Fields(sc.Bytes())
		for i := 0; i < n; i++ {
			a[i] = toInt(tmp[i])
		}
	}
	if sc.Scan() {
		tmp := bytes.Fields(sc.Bytes())
		for i := 0; i < 4; i++ {
			op[i] = toInt(tmp[i])
		}
	}

	for i := 0; i < 4; i++ {
		for j := 0; j < op[i]; j++ {
			ops[idx] = i + 1
			idx++
		}
	}
	dfs(0, 1, a[0])

	fmt.Println(max64, min64)
}
func dfs(cnt int, idx int, num int) {
	result := 0
	if cnt == n-1 {
		if num < min64 {
			min64 = num
		}
		if num > max64 {
			max64 = num
		}
	} else {
		for i := 0; i < n-1; i++ {
			if check[i] == false {
				switch ops[i] {
				case 1:
					result = num + a[idx]
				case 2:
					result = num - a[idx]
				case 3:
					result = num * a[idx]
				case 4:
					result = num / a[idx]
				}
				check[i] = true
				dfs(cnt+1, idx+1, result)
				check[i] = false
				result = a[0]
			}
		}
	}
}
func toInt(bytes []byte) int {
	n, f := 0, 1
	for _, v := range bytes {
		if v == '-' {
			f = -1
		}
		n *= 10
		n += int(v - '0')
	}
	n = n * f
	return n
}

func nojam10799() {
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
						tail = k
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

}

func main() {
	nojam10799()
}
