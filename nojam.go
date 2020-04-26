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

func main() {
	nojam11650()
}
