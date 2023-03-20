package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	var j int
	for j = 0; j < 10; j++ {
		fmt.Println(j)
	}

	sum := 1
	for sum < 20 {
		sum += sum
	}

	// Аналог while
	for sum <= 25 {
		sum += 1
		fmt.Println(sum)
	}

	// Бесконечный цикл
	for {
		fmt.Println("Stop me please!")
	}
}
