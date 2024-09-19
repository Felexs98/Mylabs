package main

import "fmt"

// Функция для вычисления суммы и разности двух чисел с плавающей запятой
func calculate(a, b float64) (float64, float64) {
	sum := a + b        // Сумма
	difference := a - b // Разность
	return sum, difference
}

func main() {
	// Пример использования функции
	num1 := 10.5
	num2 := 4.3

	// Вызов функции и получение результатов
	sum, difference := calculate(num1, num2)

	// Вывод результатов
	fmt.Println("Сумма:", sum)
	fmt.Println("Разность:", difference)
}
