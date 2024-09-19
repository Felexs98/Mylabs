package main

import "fmt"

// Функция для вычисления среднего значения трёх целых чисел
func average(a, b, c int) int {
	sum := a + b + c // Сумма трёх чисел
	avg := sum / 3.0 //
	return avg
}

func main() {
	// Три целых числа
	num1 := 12
	num2 := 8
	num3 := 10

	// Вычисляем среднее значение
	avg := average(num1, num2, num3)

	// Выводим результат
	fmt.Println("Среднее значение:", avg)
}
