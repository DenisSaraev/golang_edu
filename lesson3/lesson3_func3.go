package main

import "fmt"

func main() {
	dollar := 30

	getDollarValue := func() int {
		return dollar
	}

	fmt.Println(getDollarValue())
	dollar = 70

	fmt.Println(getDollarValue())
}
