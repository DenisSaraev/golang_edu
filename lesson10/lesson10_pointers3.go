package main

import "fmt"

func main() {
	/*
		В данном примере при помощь указателей мы отличаем
		дефолтные значения от нулевых.
		wallet1 - дефолт, wallet2 - задан.
	*/
	var wallet1 *int
	// дефолтное значение = nil, т.к. значение не задано
	fmt.Println(hasWallet(wallet1))

	wallet2 := 0
	fmt.Println(hasWallet(&wallet2))

	wallet3 := 100
	fmt.Println(hasWallet(&wallet3))
}

func hasWallet(money *int) bool {
	return money != nil
}
