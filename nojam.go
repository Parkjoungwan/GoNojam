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

func nojam4999() {
	var me string
	var doc string
	fmt.Scan(&me)
	fmt.Scan(&doc)
	var mecount int
	var doccount int
	for _, i := range me {
		if string(i) == "a" {
			mecount++
		}
	}
	for _, j := range doc {
		if string(j) == "a" {
			doccount++
		}
	}

	if mecount >= doccount {
		fmt.Println("go")
		return
	}
	fmt.Println("no")
}

func nojam10988() {
	var palen string
	fmt.Scan(&palen)
	var count int
	for i := 0; i < len(palen)/2; i++ {
		if palen[i] == palen[len(palen)-i-1] {
			count++
		}
	}
	if count == len(palen) {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}

}
func main() {
	nojam10988()
}
