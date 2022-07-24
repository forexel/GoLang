package SimpleNum

import (
	"fmt"
	"math"
	"errors"
)

var ErrorNotFound = errors.New("Простых чисел не найдено")

func SimpleNum (num int) (int,error) {
	var x, y, n int
	nsqrt := math.Sqrt(float64(num)) //записываем в переменную квадратный корень конечного числа
	is_prime := make([]bool, num) //создаем слайс длинной в заданное число из булевых перменных
if num > 3 { //если число больше трех то выполняем функцию
	for x = 1; float64(x) <= nsqrt; x++ { //запускаем цикл пометки чисел простыми
		for y = 1; float64(y) <= nsqrt; y++ { 
			n = 4*(x*x) + y*y
			if n <= num && (n%12 == 1 || n%12 == 5) {
				is_prime[n] = !is_prime[n]
			}
			n = 3*(x*x) + y*y
			if n <= num && n%12 == 7 {
				is_prime[n] = !is_prime[n]
			}
			n = 3*(x*x) - y*y
			if x > y && n <= num && n%12 == 11 {
				is_prime[n] = !is_prime[n]
			}
		}
	}

	for n = 5; float64(n) <= nsqrt; n++ {
		if is_prime[n] {
			for y = n * n; y < num; y += n * n {
				is_prime[y] = false
			}
		}
	}

	is_prime[2] = true //на число 2 вешаем маркет простого
	is_prime[3] = true //на число 2 вешаем маркет простого

	primes := make([]int, 0, 1270606) //заполняем слайс числами
	for x = 0; x < len(is_prime)-1; x++ {
		if is_prime[x] {
			primes = append(primes, x)
		}
	}
// в слайсе прайм все простые числа до N
	// let's print them
	for _, x := range primes {
		fmt.Println(x) //выводим все подходящие значения ряда
	}
}else if num == 2 { //если введено 2
	return 2,nil
}else if num == 3 { //если введено 3
	return 3,nil
}else { //в остальных случая пишем что простых чисел нет
	return 0,ErrorNotFound
}
}