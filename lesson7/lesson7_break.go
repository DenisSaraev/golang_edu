package main

import "fmt"

func main() {
	for i := 0; i <= 20; i++ {
		if i%2 == 1 {
			continue
		}
		fmt.Println(i)
	}

	for i := 0; i <= 20; i++ {
		if i >= 10 {
			break
		}
		fmt.Println(i)
	}

}
