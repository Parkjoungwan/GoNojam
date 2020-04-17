package main

import "fmt"

func nojam11050() {
	var a int
	var b int

	fmt.Scan(&a)
	fmt.Scan(&b)
	var FacA int = 1
	var FacB int = 1
	var FacC int = 1
	for i := 1; i <= a; i++ {
		FacA *= i
	}
	for j := 1; j <= b; j++ {
		FacB *= j
	}
	for k := 1; k <= (a - b); k++ {
		FacC *= k
	}
	var resultA int
	resultA = FacA / (FacB * FacC)
	fmt.Println(resultA)
}

func nojam2743() {
	var a string
	fmt.Scan(&a)
	fmt.Println(len(a))
	for _, s := range a {
		fmt.Println(string(s)) // string은 C랑 다르게 마지막에 '\0'가 없다.
	}
}

func main() {
	nojam2743()
}
