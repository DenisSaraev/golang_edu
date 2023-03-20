package main

import "fmt"

/*
Указатель - тип данных, который в кач-ве значения хранит
адрес ячейки памяти значения, либо другого указателя
(может быть nil)
*Type
*/

func main() {
	// nil указатель
	var intPointer *int
	// nil нельзя раименовать, вызовет panic
	fmt.Printf("Тип переменной nil: %T\n"+
		"Значение переменной nil: %#v \n\n",
		intPointer, intPointer)

	// не nil указатель
	var a int64 = 7
	fmt.Printf("Тип переменной а: %T \n"+
		"Значение переменной а: %#v \n\n",
		a, a)

	var pointerA = &a
	fmt.Printf("Тип указателя на а: %T \n"+
		"Значение указателя на а: %#v \n"+
		"Разименованный указатель на а: %#v \n\n",
		pointerA, pointerA, *pointerA)

	// второй способ создание не nil указателя
	// new создаёт не nil указатель
	// значение будет равно 0, т.к. это дэфолт для float
	var newPointer = new(float32)
	fmt.Printf("Тип указателя new: %T \n"+
		"Значение указателя new: %#v \n"+
		"Разименованный указатель new: %#v \n\n",
		newPointer, newPointer, *newPointer)
	// изменение переменной через разименованный указатель
	*newPointer = 3
	fmt.Printf("Тип указателя new: %T \n"+
		"Значение указателя new: %#v \n"+
		"Разименованный указатель new: %#v \n\n",
		newPointer, newPointer, *newPointer)
}
