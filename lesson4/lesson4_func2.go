package main

import "fmt"

/*
Функции названные с заглавной буквы - экспортируемые.
Их можно вызвать из другой программы, типа "fmt.Println".
Функции с малой буквы вызвать из другой программы нельзя.
*/

func main() {
	first, second := 1, 2

	var multiplier func(x, y int) int

	multiplier = func(x, y int) int { return x * y }
	fmt.Println(multiplier(first, second))

	divider := func(x, y int) int { return x / y }
	fmt.Println(divider(second, first))

	sumFunc := func(x, y int) int {
		return x + y
	}

	fmt.Println(calculate(first, second, sumFunc))
}

func calculate(x, y int, action func(x, y int) int) int {
	return action(x, y)
}
