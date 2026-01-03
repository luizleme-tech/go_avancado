package main

import "fmt"

type Instument interface {
	Sound() string
}

type Guitar struct{}

func (Guitar) Sound() string {
	return "Guitar!"
}

type Piano struct{}

func (Piano) Sound() string {
	return "Piano!"
}

func takeInstument(i Instument) {
	switch t := i.(type) {
	case *Guitar:
		fmt.Print(t.Sound())
	case *Piano:
		fmt.Println(t.Sound())
	}
}

func main() {
	takeInstument(&Guitar{})
}
