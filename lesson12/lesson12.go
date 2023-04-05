package main

import (
	"fmt"
	"time"
)

func main() {

}

type Square struct {
	Side int
}

func (s Square) Perimetr() {
	fmt.Printf("%T, %#v \n", s, s)
	fmt.Printf("%T, %#v \n", s, s)
}

func (s Square) WrongSc