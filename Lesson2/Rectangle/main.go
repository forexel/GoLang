package main

import (
	"fmt"
)

func main() {
	var a float32
	var b float32
	fmt.Println("Введите Длину")
	fmt.Scanln(&a)
	fmt.Println("Введите Ширину")
	fmt.Scanln(&b)

	if (a >= 0) && (b >= 0) { 
		fmt.Println(a * b)
	} else {
		fmt.Println("Я не могу посчитать площадь с такими значениями")
	}
}
