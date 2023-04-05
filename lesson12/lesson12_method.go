package main

import (
	"fmt"
)

func main() {
	definition()
	rules()
}

type Square struct {
	Side int
}

func (s Square) Perimetr() {
	/*
	   Значение в скобках - велью ресивер:
	   s - переменная ресивера, доступная внутри функции
	   Square - структура внутри которой доступна функция Perimetr
	*/
	fmt.Printf("%T, %#v \n", s, s)
	fmt.Printf("Периметр фигуры: %d \n", s.Side*4)
}

func (s *Square) Scale(multiplier int) {
	// Поинтер ресивер
	fmt.Printf("%T, %#v \n", s, s)
	s.Side *= multiplier
}

func (s Square) WrongScale(multiplier int) {
	// Отличается от Scale тем, что это велью ресивер.
	// Из-за этого из функции нельзя повлиять на внешнее значение.
	fmt.Printf("%T, %#v \n", s, s)
	s.Side *= multiplier
	fmt.Printf("%T, %#v \n", s, s)
}

func definition() {
	square := Square{Side: 4}
	pSquare := &square

	square2 := Square{Side: 2}

	square.Perimetr()
	square2.Perimetr()

	/*
		Пример синтаксического сахара.
		Не смотря на то, что функция Perimetr ссылается НЕ на поинтер,
		нам ничего не мешает передать ей поинтер (указатель).
		И наоборот с функцией Scale, которая ожидает в кач-ве ресивера
		поинтер, но мы передаём ей обычное значение.
		Golang автоматически !!!делает обратную подстановку!!!.
	*/

	pSquare.Scale(2)
	square.Scale(2)
	pSquare.Perimetr()

	// WrongScale - велью ресивер. Не смог увеличить Side.
	// Поэтому не верное значение Perimetr.
	square.WrongScale(2)
	square.Perimetr()
}

func rules() {
	/*
		Поинтер или велью:
		а) Если нужно менять значение - применять поинтер ресивер.
		б) Поинтер - строка с адресом и меньше весит, чем велью.
		в) Много указателей перегружают GarbageCollector (GC).
		г) Пакет sync использовать только по указателю.

		Ограничения методов.
		Мы можем добавить метод для типа Т или Т* при условиях:
		а) Т находится в данном пакете.
		б) Т не является указателем.
		в) Т - конкретный, не интерфейсный тип.
		г) Т не является builtin (встроенным) типом.
		д) Т является объявленным типом.
	*/
}
