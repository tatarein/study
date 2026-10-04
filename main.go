package main

import "fmt"

func main() {
	fmt.Println("Расписание на неделю:")

	// Переменная day начинается с 1; цикл работает, пока day <= 7; после каждого круга day увеличивается на 1 (day++)
	for day := 1; day <= 7; day++ {
		if day <= 5 {
			fmt.Println("День", day, "- Смена + 2 часа учебы")
		} else {
			fmt.Println("День", day, "- Выходной! 4 часа полного погружения в код")
		}
	}
}
