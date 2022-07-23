package main

import "fmt"

// описываю функцию фибонначи
func fib(n uint, hash map[uint]uint) uint {
	if n < 2 { //если индекс = 0, то выводим 0
		return n
	}
	if val, ok := hash[n]; ok {
		return val
	}

	hash[n] = fib(n-1, hash) + fib(n-2, hash)

	return hash[n]
}

func main() {
	var n uint
	hash := make(map[uint]uint)
	fmt.Println("Введите порядковый номер числа Фибоначчи:")
	fmt.Scanln(&n)
	if n < 1 {
		fmt.Println("Некорректный номер элемента")
	} else {
		fmt.Println(fib((n - 1), hash)) //ищем значение элемента с индексом порядковый номер-1
	}
}
