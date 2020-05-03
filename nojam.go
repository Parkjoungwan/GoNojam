package main

import (
	"bufio"
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

func nojam11866() {

}
func main() {

}
