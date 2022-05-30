package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Println("Введите трехзначное целое число")
	fmt.Scanln(&a)
	b := a / 100
	c := a/10 - b*10
	d := a - b*100 - c*10
	if (100 <= a) && (a <= 999) {
		fmt.Println("Число содержит", b, getFirst(b, "сотню", "сотни", "сотен"), c, getFirst(c, "десяток", "десятка", "десятков"), d, getFirst(d, "единицу", "единицы", "единиц"))
	} else {
		fmt.Println("Введено не трехначное целое число")
	}
}

func getFirst(number int, word1 string, word2 string, word3 string) string {
	var n int
	n = number
	if n == 1 {
		return word1
	} else if (n >= 2) && (n <= 4) { 
		return word2
	} else {
		return word3
	}
}