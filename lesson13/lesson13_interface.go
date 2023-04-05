package main

import "fmt"

type Runner interface {
	Run() string
}

type Swimmer interface {
	Swim() string
}

type Flyer interface {
	Fly() string
}

type Ducker interface {
	Runner
	Swimmer
	Flyer
}

type Human struct {
	name string
}

func main() {
	interfaceValues()
}

func interfaceValues() {
	var runner Runner
	fmt.Printf("Type: %T Value: %#v\n", runner, runner)

	if runner == nil {
		fmt.Println("Runner is nil")
	}
}
