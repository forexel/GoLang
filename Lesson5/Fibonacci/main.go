package main

import "fmt"

// описываю функцию фибонначи
func fib(n uint) uint {
	if n == 0 { //если индекс = 0, то выводим 0
		return 0
	} else if n == 1 { //Если индекс равен 1, то выводим 1
		return 1
	} else {
		return fib(n-1) + fib(n-2) //в противном случае рассчитываем сумму двух чисел фиббоначи
	}
}

func main() {
	var n uint
	fmt.Println("Введите порядковый номер числа Фибоначчи:")
	fmt.Scanln(&n)
	if n < 1 {
		fmt.Println("Некорректный номер элемента")
	} else {
		fmt.Println(fib(n - 1)) //ищем значение элемента с индексом порядковый номер-1
	}
}
