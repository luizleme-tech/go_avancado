package main

//func main() {
//
//	x := 10
//	take(&x)
//	fmt.Println(x)
//	// aqui exemplo inicial de ponteiro
//	//p := &x
//	//fmt.Println(p, *p)
//}

//func main() {
//	x := create()
//	fmt.Println(*x)
//}

//func main() {
//	// erro de dereferencia em runtime
//	var x *int
//	take(x)
//	fmt.Println(x)
//}

func create() *int {
	x := 10
	return &x
}

func take(x *int) {
	*x = 100
}

// p = 0x123
// p -> 0x123 = ??

// por ponteiro é o byRef sem ponteiro é o byVal
// por byRef isso e possivel alterar o valor da variavel
