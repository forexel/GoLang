package main

import "fmt"

func main() {
	var sl []int
	var n int
	fmt.Println("Введите количество чисел в ряду:")
	if m, err := Scan(&n); m != 1 {
		panic(err)
	}
	fmt.Println("Введите", n, "чисел")
	sl = make([]int, n)
	ReadN(sl, 0, n)
	Sort(sl)
	fmt.Println(sl)
}

func Sort(sl []int) {
	for i := 2; i < len(sl); i++ {
		x := sl[i]
		j := i
		for ; j >= 1 && sl[j-1] > x; j-- {
			sl[j] = sl[j-1]
		}
		sl[j] = x
	}
}

func input(x []int, err error) []int {
	if err != nil {
		return x
	}
	var d int
	n, err := fmt.Scanf("%d", &d)
	if n == 1 {
		x = append(x, d)
	}
	return input(x, err)
}

func ReadN(all []int, i, n int) {
	if n == 0 {
		return
	}
	if m, err := Scan(&all[i]); m != 1 {
		panic(err)
	}
	ReadN(all, i+1, n-1)
}

func Scan(a *int) (int, error) {
	return fmt.Scan(a)
}
