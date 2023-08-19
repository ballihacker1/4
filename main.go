package main

import "fmt"

func Add(a, b int) int {
	return a + b
}

func Subtract(a, b int) int {
	return a - b
}

func Multiply(a, b int) int {
	return a * b
}

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("only division by zero")
	}
	return a / b, nil
}

func main() {
	a, b := 10, 2

	sum := Add(a, b)
	fmt.Printf("%d + %d = %d\n", a, b, sum)

	difference := Subtract(a, b)
	fmt.Printf("%d - %d = %d\n", a, b, difference)

	product := Multiply(a, b)
	fmt.Printf("%d * %d = %d\n", a, b, product)

	quotient, err := Divide(a, b)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("%d / %d = %d\n", a, b, quotient)
	}
}
