package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

func nojam9020() {
	wr := bufio.NewWriter(os.Stdout)
	sc := bufio.NewScanner(os.Stdin)
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
	wr := bufio.NewWriter(os.Stdout)
	sc := bufio.NewScanner(os.Stdin)
	defer wr.Flush()
	var T int
	if sc.Scan() {
		T, _ = strconv.Atoi(sc.Text())
	}
	Num := make([]int, T)
	for i := 0; i < T; i++ {
		if sc.Scan() {
			Num[i], _ = strconv.Atoi(sc.Text())
		}
	}
	sort.Sort(sort.IntSlice(Num))
	var sum int
	var sub int
	var mul int
	var div int
	if sc.Scan() {
		sum, _ = strconv.Atoi(sc.Text())
	}
	if sc.Scan() {
		sub, _ = strconv.Atoi(sc.Text())
	}
	if sc.Scan() {
		mul, _ = strconv.Atoi(sc.Text())
	}
	if sc.Scan() {
		div, _ = strconv.Atoi(sc.Text())
	}
	var max int
	var min int
	for i:=0;i<mul;i+=2{
		max=Num[T-i]*Num[T-(i+1)]
		min=Num[i]*Num[i+1]
	}
	for i:=0;i<sum;i++{
		
	}
}

func main() {
	nojam9020()
}
