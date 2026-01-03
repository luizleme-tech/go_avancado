package main

import "fmt"

//func main() {
//	arr := [5]int{1, 2, 3, 4, 5}
//	slice := arr[:2:2]
//	fmt.Println(slice, cap(slice))
//}

//func main() {
//	slice := []int{1, 2, 3, 4}
//	foo(slice)
//}
//
//func foo(slice []int) {
//	_ = slice[3] // bounds check
//	fmt.Println(slice[0])
//	fmt.Println(slice[1])
//	fmt.Println(slice[2])
//	fmt.Println(slice[3])
//}

func main() {
	slice := []int{1, 2, 3, 4}
	foo(slice)
	fmt.Println(slice)
}

func foo(slice []int) {
	slice[0] = 123 // aqui se altera o slice pois ele é passado por referencia, para ser por valor deveria ser um array
}
