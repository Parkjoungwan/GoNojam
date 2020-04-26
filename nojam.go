package main

import (
	"fmt"
	"sort"
)

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

func nojam3047() {
	Num := make([]int, 3)
	for i := 0; i < 3; i++ {
		fmt.Scan(&Num[i])
	}
	sort.Sort(sort.IntSlice(Num))
	A := Num[0]
	B := Num[1]
	C := Num[2]
	var str string
	fmt.Scan(&str)
	if str[0] == 'A' {
		fmt.Print(A, " ")
	} else if str[0] == 'B' {
		fmt.Print(B, " ")
	} else {
		fmt.Print(C, " ")
	}
	if str[1] == 'A' {
		fmt.Print(A, " ")
	} else if str[1] == 'B' {
		fmt.Print(B, " ")
	} else {
		fmt.Print(C, " ")
	}
	if str[2] == 'A' {
		fmt.Print(A, " ")
	} else if str[2] == 'B' {
		fmt.Print(B, " ")
	} else {
		fmt.Print(C, " ")
	}
}
func main() {
	nojam3047()
}
