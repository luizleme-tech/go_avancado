package main

type Animal interface {
	Sound() string
}

type Dog struct{}

func (Dog) Sound() string {
	return "Woof!"
}

func whatDoesThisAnimalSay(a Animal) {
	println(a.Sound())
}

func mainInterface() {
	dog := Dog{}
	whatDoesThisAnimalSay(dog)
}

// INTERFACE VAZIA
//func foo(a interface{})
//ou
//func foo(a any)
