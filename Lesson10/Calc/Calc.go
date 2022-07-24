package Calc

import (
	"errors"
	"fmt"
	"math"
)

var ErrorFactNegative = errors.New("Невозможно посчитать факториал отрицательного числа")
var ErrorFactFloatNumber = errors.New("Умею считать факториал такого числа")
var ErrorUnknownOperation = errors.New("Неизвестная операция")
var ErrorDivideZero = errors.New("Делить на ноль нельзя")
var ErrorRoot = errors.New("Не умею извлекать корень")

func Calc(a float64, b float64, op string) (float64, error) {
	var result float64
	if op == "!" {
		if a < 0 {
			return 0, ErrorFactNegative
		} else if a == 0 {
			return 1, nil //Если число 0, то факториал равен 1
		} else if int(a*100)%100 == 0 { //считаем факториал
			result = 1
			for i := 1; i < int(a+1); i++ { //запускаем счетчик
				result = result * float64(i)
			}
			return result, nil // выводим число
		} else {
			return 0, ErrorFactFloatNumber
		}
	} else {
		switch {
		case op == "+":
			{ // если оператор плюс
				result = a + b
				return result, nil
			}
		case op == "-":
			{ // если оператор минус
				result = a - b
				return result, nil
			}
		case op == "^":
			{ //если оператор степень
				if b < 0 { //если степень меньше нуля
					return 0, ErrorRoot
				} else {
					result = math.Pow(a, b)
					return result, nil
				}
			}
		case op == "*":
			{ //если оператор умножить
				result = a * b
				return result, nil
			}
		case op == "/":
			{ //если оператор делить
				if b == 0 {
					return 0, ErrorDivideZero
				} else {
					res := float64(a / b)
					return res, nil
				}
			}
		default:
			fmt.Println("Неизвестная операция") //если введен какой-то неизвестный символ - выводим ошибку
			return 0, ErrorUnknownOperation
		}
	}
}
