package solved

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

var (
	sc = bufio.NewScanner(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)
)

func Nojam11050() {
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

func Nojam2743() {
	var a string
	fmt.Scan(&a)
	fmt.Println(len(a))
	for _, s := range a {
		fmt.Println(string(s)) // string은 C랑 다르게 마지막에 '\0'가 없다.
	}
}

func Nojam4999() {
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

func Nojam10988() {
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
func Nojam1002() {
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

func Nojam2941() {
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

func Nojam11726() {
	var ga [1001]int
	var N int
	fmt.Scan(&N)
	ga[1] = 1
	ga[2] = 2
	for i := 3; i <= N; i++ {
		ga[i] = ga[i-1]%10007 + ga[i-2]%10007
	}
	fmt.Println(ga[N])
}

var ans int

func Nojam9663() {
	var N int
	fmt.Scan(&N)
	for i := 1; i <= N; i++ {
		var Col []int = make([]int, 15)
		Col[1] = i
		Nojam9633dfs(1, N, Col)
	}
	fmt.Println(ans)
}

func Nojam9633dfs(Row int, N int, Col []int) {
	if Row == N {
		ans++
	} else {
		for i := 1; i <= N; i++ {
			Col[Row+1] = i
			if Nojam9633Possible(Row+1, Col) {
				Nojam9633dfs(Row+1, N, Col)
			} else {
				Col[Row+1] = 0
			}
		}
	}
}

func Nojam9633Possible(C int, Col []int) bool {
	for i := 1; i < C; i++ {
		if Col[i] == Col[C] {
			return false
		}
		if math.Abs(float64(Col[i]-Col[C])) == math.Abs(float64(i-C)) {
			return false
		}
	}
	return true
}

func Nojam11650() {
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

func Nojam16212() {
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

func Nojam1059() {
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

func Nojam1041() {
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

func Nojam3047() {
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

func Nojam2309() {
	Height := make([]int, 9)
	var all int
	var i, j int
	var fake1, fake2 int
	for i := 0; i < 9; i++ {
		fmt.Scan(&Height[i])
	}
	sort.Sort(sort.IntSlice(Height))
	for i = 0; i < 9; i++ {
		all += Height[i]
	}
	for i = 0; i < 9; i++ {
		for j = 0; j < 9; j++ {
			if all-Height[i]-Height[j] == 100 {
				fake1 = i
				fake2 = j
			}
		}
	}
	for k := 0; k < 9; k++ {
		if fake1 == k || fake2 == k {
		} else {
			fmt.Println(Height[k])
		}
	}
}

func Nojam1699() {
	var number [100001]int
	var N int
	fmt.Scan(&N)
	for i := 1; i <= N; i++ {
		for j := 1; j*j <= i; j++ {
			if number[i] > number[i-j*j]+1 || number[i] == 0 {
				number[i] = number[i-j*j] + 1
			}
		}
	}
	fmt.Println(number[N])
}

func Nojam1010() {
	var N, M int
	var T int
	fmt.Scan(&T)
	var DP [31][31]int
	for i := 1; i <= 30; i++ {
		DP[1][i] = i
	}
	for i := 2; i <= 30; i++ {
		for j := i; j <= 30; j++ {
			for k := j; k >= i; k-- {
				DP[i][j] = DP[i][j] + DP[i-1][k-1]
			}
		}
	}
	for i := 0; i < T; i++ {
		fmt.Scan(&N, &M)
		fmt.Println(DP[N][M])
	}
}

func Nojam11399() {
	var N int
	fmt.Scan(&N)
	Per := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&Per[i])
	}
	sort.Sort(sort.IntSlice(Per))
	var ans int
	for i := 0; i < N; i++ {
		for j := 0; j <= i; j++ {
			ans += Per[j]
		}
	}
	fmt.Println(ans)
}

func Nojam2875() {
	var M int
	var F int
	var K int
	fmt.Scan(&F, &M, &K)
	for K != 0 {
		if F > M*2 {
			F -= 1
			K -= 1
		} else {
			M -= 1
			K -= 1
		}
	}
	var count int
	for M != 0 {
		F -= 2
		if F < 0 {
			break
		}
		M -= 1
		count++
	}
	fmt.Println(count)

}
func Nojam10610() {
	var s string
	var sum int64
	sc := bufio.NewScanner(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)
	digit := make([]int, 10)
	const maxCapacity = 512 * 1024
	buf := make([]byte, maxCapacity)
	sc.Buffer(buf, maxCapacity)
	defer wr.Flush()
	if sc.Scan() {
		s = sc.Text()
	}
	for i := 0; i < len(s); i++ {
		var tmp = int64(s[i] - '0')
		digit[tmp]++
		sum += tmp
	}
	if digit[0] == 0 || sum%3 != 0 {
		fmt.Fprintf(wr, "-1")
	} else {
		for i := 9; i >= 0; i-- {
			for j := 0; j < digit[i]; j++ {
				fmt.Fprintf(wr, "%d", i)
			}
		}
	}
}
func Nojam1929() {
	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()
	var N, M int
	fmt.Scan(&N, &M)
	var PN [1000001]int
	PN[1] = 1
	for i := 2; i <= M; i++ {
		if PN[i] == 0 {
			for j := i * i; j <= M; j += i {
				PN[j] = 1
			}
		}
	}
	for i := N; i <= M; i++ {
		if PN[i] != 1 {
			fmt.Fprintln(wr, i) //이거 출력이 안될지도?
		}
	}
}

func Nojam9020() {
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
func Nojam14888() {
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

func Nojam10799() {
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
