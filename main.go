package main

import "fmt"

func main() {
	sneakerPrice := 15000
	fmt.Println("Система проверки баланса запущена. (Для выключения введите 0)")

	// Этот цикл будет крутиться вечно, пока мы его не остановим
	for {
		var budget int
		fmt.Print("\nВведите ваш текущий бюджет: ")
		fmt.Scan(&budget)

		// Секретная кнопка выхода
		if budget == 0 {
			fmt.Println("Сервер остановлен. До свидания!")
			break // Эта команда мгновенно ломает цикл
		}

		// Твоя рабочая логика проверок
		if budget >= sneakerPrice {
			fmt.Println("Покупка одобрена! Остаток:", budget-sneakerPrice)
		} else {
			fmt.Println("Недостаточно средств. Вам не хватает:", sneakerPrice-budget)
		}
	}
}
