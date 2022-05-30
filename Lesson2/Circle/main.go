package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64
	var b float64
	var c float64
	var d float64
	var l float64 // назначаю всем переменным тип флоат, потому что будем делить на дробное число
	fmt.Println("Введите Площадь круга")
	fmt.Scanln(&a) // читаем значение площади
	if a >= 0 {    //проверяем, что введена площадь больше или равная нулю
		b = a / math.Pi  //считаем квадрат радиуса как площадь/пи
		c = math.Sqrt(b) //считаем радиус как корень из квадрата радиуса
		d = 2 * c        //считаем диаметр как удвоенный радиус
		l = d * math.Pi  // считаем длинку окружности как диаметр умноженный на пи
		fmt.Println("Диаметр круга равен", d)
		fmt.Println("Длинна окружности равна", l)
	} else { //обрабатываем противный случай
		fmt.Println("Я не могу посчитать диаметр и площадь с таким значение площади") 
	}
}
