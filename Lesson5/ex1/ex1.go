package main

import "fmt"

type Person struct {
	FirstName string
	LanstName string
	Age       uint8
}

func main() {
	anna := Person{
		FirstName: "Anna",
		LanstName: "Ivanova",
		Age:       28,
	}
	fmt.Printf("%+v\n", anna.Age)
}
