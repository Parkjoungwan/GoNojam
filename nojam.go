package main

import (
	"fmt"
	"math"
)

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
	if count == len(palen)/2 {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}
}
func nojam1002() {
	var x1, x2, y1, y2, r1, r2 float64
	var count int
	fmt.Scan(&count)
	for i := 0; i < count; i++ {
		fmt.Scan(&x1, &y1, &r1, &x2, &y2, &r2)
		bewTwo := (x1-x2)*(x1-x2) + (y2-y1)*(y2-y1)
		bewTwo = math.Sqrt(bewTwo)
		if x1 == x2 && y1 == y2 {
			if r1 == r2 {
				fmt.Println(-1)
			} else {
				fmt.Println(0)
			}
		} else if bewTwo == r1+r2 {
			fmt.Println(1)
		} else if bewTwo < r1+r2 {
			if math.Abs(r2-r1) > bewTwo {
				fmt.Println(0)
			} else if math.Abs(r2-r1) == bewTwo {
				fmt.Println(1)
			} else {
				fmt.Println(2)
			}
		} else if bewTwo > r1+r2 {
			fmt.Println(0)
		}
	}
}

func nojam2941() {
	var Cro string
	var count int
	var count2 int
	var dzcount int
	fmt.Scan(&Cro)
	for i := 0; i < len(Cro); i++ {
		if string(Cro[i]) == "c" {
			if i == len(Cro)-1 {
				break
			}
			if string(Cro[i+1]) == "=" || string(Cro[i+1]) == "-" {
				count++
				count2 += 2
			}
		} else if string(Cro[i]) == "d" {
			if i == len(Cro)-1 {
				break
			}
			if i == len(Cro)-1 {
				break
			}
			if string(Cro[i+1]) == "-" {
				count++
				count2 += 2
			} else if string(Cro[i+1]) == "z" {
				if i+1 == len(Cro)-1 {
					break
				}
				if string(Cro[i+2]) == "=" {
					count++
					count2++
					dzcount++
				}
			}
		} else if string(Cro[i]) == "l" {
			if i == len(Cro)-1 {
				break
			}
			if string(Cro[i+1]) == "j" {
				count++
				count2 += 2
			}
		} else if string(Cro[i]) == "n" {
			if i == len(Cro)-1 {
				break
			}
			if string(Cro[i+1]) == "j" {
				count++
				count2 += 2
			}
		} else if string(Cro[i]) == "s" {
			if i == len(Cro)-1 {
				break
			}
			if string(Cro[i+1]) == "=" {
				count++
				count2 += 2
			}
		} else if string(Cro[i]) == "z" {
			if i == len(Cro)-1 {
				break
			}
			if string(Cro[i+1]) == "=" {
				count++
				count2 += 2
			}
		}
	}
	fmt.Println((len(Cro) - count2) + (count - dzcount))
}

func main() {
	nojam2941()
}
