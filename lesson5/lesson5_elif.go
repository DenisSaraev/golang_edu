package main

import "fmt"

/*
&& "и"
 true && true = true
 true && false = false
 false && true = false
 false && false = false

|| "или"
 true || true = true
 true || false = true
 false || true = true
 false || false = false

! "не"
 !true = false
 !false = true
*/

func main() {
	age := 14

	if age < 18 {
		fmt.Println("You are too young")
	}

	if isChild := isChildren(age); isChild == true {
		fmt.Println("You are child")
	}

	age = 20

	if age < 18 {
		fmt.Println("You are too young")
	} else {
		fmt.Println("You are an adult")
	}

	age = 14

	if age >= 7 && age <= 18 {
		fmt.Println("You are in school")
	}

	age = 14

	if age == 14 || age == 21 || age == 40 {
		fmt.Println("You have to get new passport")
	}

	age = 20

	if !isChildren(age) {
		fmt.Println("You are not young")
	}

}

func isChildren(age int) bool {
	return age < 18
}
