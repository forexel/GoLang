package main

import "fmt"

func plusAndMulti(arr ...int) (sum int, mult int) {
	mult = 1

	for _, val := range arr {
		sum += val
		mult *= val
	}
	return
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	sum, mult := plusAndMulti(arr...)

	fmt.Println(sum, mult)
}
