package main

import "fmt"

func main() {
	i := 0
	fmt.Println(i)

	//Одно и тоже разным синтаксисом
	//Инкремент
	i = i + 1
	fmt.Println(i)
	i += 1
	fmt.Println(i)
	i++
	fmt.Println(i)

	//Декремент
	i = i - 1
	fmt.Println(i)
	i -= 1
	fmt.Println(i)
	i--
	fmt.Println(i)
}
