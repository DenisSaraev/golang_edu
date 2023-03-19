package main

import "fmt"

func main() {
	Greet()
	PersonGreet("Denis")
	FioGreet("Denis", "Saraev")

	first, second := 2, 5
	summary := Sum(first, second)
	fmt.Println(summary)

	sum, mult := SumAndMultiply(first, second)
	fmt.Println(sum, mult)

	_, mult64 := NamedSumAndMultiply(first, second)
	// Первую переменную отправили в null
	fmt.Println(mult64)
}

func Greet() {
	fmt.Println("Hello world")
}

func PersonGreet(name string) {
	fmt.Printf("Hi %s\n", name)
}

func FioGreet(name, surname string) {
	fmt.Printf("Hi %s %s\n", name, surname)
}

func Sum(first, second int) int {
	sum := first + second
	return sum
}

func SumAndMultiply(first, second int) (int, int) {
	return first + second, first * second
}

func NamedSumAndMultiply(first, second int) (sum int64, multiply int64) {
	sum = int64(first + second)
	multiply = int64(first) * int64(second)
	return // или return sum, multiply
}
