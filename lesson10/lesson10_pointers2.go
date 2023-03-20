package main

import "fmt"

func main() {
	num := 3

	/*
		После возведения в квадрат значение num
		не изменилось, т.к. внутри функции создалась
		своя переменная.
		Чтобы это обойти мы могли бы сделать
		явное переопределение:
		num = square(num)
	*/
	square(num)
	fmt.Println(num)

	// т.к. в квадрат возвели значение ссылаясь
	// на ячейку памяти значение изменилось
	squarePointer(&num)
	fmt.Println(num)
}

func square(num int) {
	num *= num
}

func squarePointer(num *int) {
	*num *= *num
}
