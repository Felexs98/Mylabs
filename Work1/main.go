package main

import (
	"fmt"
	"time"
)

func main() {
	// Получаем текущее время
	currentTime := time.Now()

	// Выводим текущее время и дату
	fmt.Println("Текущее время и дата:", currentTime.Format("02-01-2006 15:04:05"))
}
