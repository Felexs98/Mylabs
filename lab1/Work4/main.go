package main

import (
	"fmt"
)

func main() {
	num1 := 15
	num2 := 4

	// Операции

	sum := num1 + num2        // Сумма
	difference := num1 - num2 // Вычет
	product := num1 * num2    // Умножение
	quotient := num1 / num2   // Деление
	remainder := num1 % num2  // Остаток от деления

	// Вывод рез

	fmt.Println("Сумма:", sum)
	fmt.Println("Вычитание:", difference)
	fmt.Println("Умножение:", product)
	fmt.Println("Деление:", quotient)
	fmt.Println("Остаток от деление:", remainder)
}
