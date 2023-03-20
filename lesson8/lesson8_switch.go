package main

import (
	"fmt"
	"math/rand" //randomizer
	"time"
)

const (
	min = 1
	max = 5
)

func main() {
	/*
		Т.к. генератор случайных чисел выдаёт псевдо-
		случайные числа, то последовательность случайных
		чисел всегда одинаковая.
		Для решения проблемы нужно передавать каждый раз
		новую отметку для отсчёта.
		В примере ниже в качестве отметки берётся текущее
		время в нано-секундах.
	*/
	rand.Seed(time.Now().UnixNano())
	// Генерируем int число от 1 до 5
	value := rand.Intn(max-min) + min

	if value == 1 {
		fmt.Println("Number is one")
	} else if value == 2 || value == 3 {
		fmt.Println("Number is two or three")
	} else if value == getFour() {
		fmt.Println("Number is four")
	} else {
		fmt.Println("Default case is shown")
	}

	switch value {
	case 1:
		fmt.Println("Number is one")
		/*
			После срабатывания case происходит выход из
			конструкции switch.
			В python пришлось бы делать break для этого
			функционала.
			Чтобы не срабатывал выход из switch и
			продолжалась обработка, нужно использовать
			fallthrough.
		*/
	case 2, 3:
		fmt.Println("Number is two or three")
	case getFour():
		fmt.Println("Number is four")
		fallthrough
	case 10:
		fmt.Println("Strange things happens")
	default:
		fmt.Println("Default case is shown")
	}

	switch {
	case value > 2:
		fmt.Printf("Value %d greater than 2 \n", value)
	case value < 2:
		fmt.Printf("Value %d less than 2 \n", value)
	default:
		fmt.Println("Value equals 2")
	}
}

func getFour() int {
	fmt.Println("getFour called")
	return 4
}
