package fib

import "errors"

var ErrorNegative = errors.New("Нет значения отрицательного элемента")

// описываю функцию фибонначи
func fibOptimized(n int, hash map[int]int) (int, error) {
	if n < 0 {
		return 0, ErrorNegative
	}
	if n < 2 { //если индекс = 0, то выводим 0
		return n, nil
	}
	if val, ok := hash[n]; ok {
		return val, nil
	}
	prev, err := fibOptimized(n-1, hash)
	if err != nil {
		return 0, err
	}
	prev2, err := fibOptimized(n-2, hash)
	if err != nil {
		return 0, err
	}
	hash[n] = prev2 + prev

	return hash[n], nil
}

func fibunOptimized(n int) (float64, error) {
	if n < 0 {
		return 0, ErrorNegative
	}
	if n == 0 { //если индекс = 0, то выводим 0
		return 0, nil
	} else if n == 1 { //Если индекс равен 1, то выводим 1
		return 1, nil
	} else {
		return (fibunOptimized(n-1) + fibunOptimized(n-2)), nil //в противном случае рассчитываем сумму двух чисел фиббоначи
	}
}
