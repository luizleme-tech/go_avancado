package main

import "fmt"

func firstmain() {
	arr := [5]int{1, 2, 3, 4, 5}
	slice := arr[1:4]
	sliceOriginal := []int{1, 2, 3, 4, 5}

	fmt.Println(slice)
	fmt.Println(sliceOriginal)
	fmt.Println(sliceOriginal, len(sliceOriginal), cap(sliceOriginal))
	// diferenca de um slice para um array, o slice e um array dinamico

	// slice length e capacity
	// length = quantidade de elementos do slice
	// capacity = capacidade do slice sem fazer uma nova alocação
}
