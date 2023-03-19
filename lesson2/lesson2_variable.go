package main

import "fmt"

/*
go run .\lesson2\lesson2_variables.go
*/

const hello3 = "You can't change me"

func main() {
	var hello string
	hello = "Hello!"
	hello2 := true
	fmt.Println(hello)
	fmt.Printf("Type: %T Value: %v\n", hello, hello)
	fmt.Println(hello2)
	fmt.Printf("Type: %T Value: %v\n", hello2, hello2)
	fmt.Println(hello3)
	fmt.Printf("Type: %T Value: %v\n", hello3, hello3)
}

/*
Типы данных:
bool (default = false)
string (default = "")

int
int8 (-128/128) int16(-32768/32768)
int32(–2147483648/2147483647) int64 (-2^63/2^63-1)

uint uint8 uint16 uint32 uint64 - больше нуля
byte (uint8)
rune (int32) - символ
float32 float64
complex64 complex128 - комплексные число (редко)
*/
