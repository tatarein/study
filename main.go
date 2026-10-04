package main

import (
	"fmt"
	"net/http"
	"strconv" // Этот пакет нужен, чтобы превратить текст из ссылки в математическое число
)

func checkBalance(w http.ResponseWriter, r *http.Request) {
	// 1. Достаем значение "budget" прямо из адресной строки браузера
	moneyFromURL := r.URL.Query().Get("budget")

	// 2. Переводим текст в число (ошибку пока игнорируем с помощью _)
	budget, _ := strconv.Atoi(moneyFromURL)
	sneakerPrice := 15000

	// 3. Наша знакомая логика, но теперь ответ улетает в браузер (используем Fprintln и букву w)
	if budget >= sneakerPrice {
		fmt.Fprintln(w, "Покупка одобрена! Остаток:", budget-sneakerPrice, "руб.")
	} else {
		fmt.Fprintln(w, "Недостаточно средств. Вам не хватает:", sneakerPrice-budget, "руб.")
	}
}

func main() {
	// Создаем новый маршрут: при переходе на /balance запустится функция checkBalance
	http.HandleFunc("/balance", checkBalance)

	fmt.Println("Сервер запущен! Проверь баланс по ссылке: http://localhost:8080/balance?budget=20000")
	http.ListenAndServe(":8080", nil)
}
