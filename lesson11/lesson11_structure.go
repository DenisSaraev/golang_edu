package main

import (
	"fmt"
	"time"
)

type OurString string
type OurInt int

type Person struct {
	Name string
	Age  int
}

func main() {
	var customString OurString
	fmt.Printf("%T %#v \n", customString, customString)
	customString = "Hello dudes"
	fmt.Printf("%T %#v \n", customString, customString)

	customInt := OurInt(5)
	fmt.Printf("%T %#v \n", customInt, customString)
	fmt.Println(int(customInt))

	var John Person
	fmt.Printf("%T %#v \n", John, John)

	John = Person{}
	fmt.Printf("%T %#v \n", John, John)

	John.Name = "John"
	John.Age = 30
	fmt.Println(John)

	Brad := Person{
		Name: "Brad",
		Age:  25,
	}
	fmt.Println(Brad)

	Vladimir := Person{"Vladimir", 40}
	fmt.Println(Vladimir)

	pVladimir := &Vladimir
	fmt.Println((*pVladimir).Age)
	fmt.Println(pVladimir.Age)

	pIvan := &Person{"Ivan", 90}
	fmt.Println(pIvan)

	unnamedStruct := struct {
		Name, LastName, BirthDate string
	}{
		Name:      "NoName",
		LastName:  "NoLastName",
		BirthDate: fmt.Sprintf("%s", time.Now()), // превращает в строку
		// BirthDate: time.Now().String(),
	}
	fmt.Println(unnamedStruct)
}
