package main

import (
	"errors"
	"fmt"
)

// Primeiro exemplo de ERRO //

// Erros em Go são valores, não lança exception (Em Go Exception é um Panic)
//func main() {
//	a := 10
//	b := 0
//	res, err := dividir(a, b)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	fmt.Println(res)
//}
//
//func dividir(a, b int) (int, error) {
//	if b == 0 {
//		return 0, errors.New("não pode dividir por zero")
//	}
//	return a / b, nil
//}

// Segundo exemplo de ERRO //

// Aqui vai ocorrer um panic
//func main() {
//	user, _ := NewUser(true)
//	user.Foo()
//}

//func main() {
//	user, err := NewUser(true)
//	if err != nil {
//		fmt.Println("Algum erro na hora de criar o usuario")
//		return
//	}
//	user.Foo()
//}

//type User struct {
//	foo string
//}

//func (u User) Foo() {
//	fmt.Println(u.foo)
//}

//func NewUser(wantErr bool) (*User, error) {
//	if wantErr {
//		return nil, errors.New("erro")
//	}
//	return &User{}, nil
//}

// Terceiro exemplo de ERRO //

//type SqrtError struct {
//	msg string
//}
//
//func (s SqrtError) Error() string { return s.msg }
//
//func raizQuadrada(x float64) (float64, error) {
//	if x < 0 {
//		return 0, SqrtError{"não existe raiz quadrada de numero negativo"}
//	}
//	resultado := math.Sqrt(x)
//	return resultado, nil
//}

//func main() {
//	x := -10
//	res, err := raizQuadrada(float64(x))
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	fmt.Println(res)
//}

// Diferenca entre Error.is e Erros.as

//var ErrNotFound = errors.New("not found")

//func foo() error { return nil }

//func foo() error { return ErrNotFound }

//func main() {
//	err := foo()
//	if err != nil && errors.Is(err, ErrNotFound) {
//		fmt.Println("deu erro not found")
//		return
//	}
//	fmt.Println("foi pra fora")
//}

//func foo() error { return SqrtError{"teste"} }

//func main() {
//	err := foo()
//	var sqrtError SqrtError
//	if err != nil && errors.As(err, &sqrtError) { // passar um ponteiro
//		fmt.Println(sqrtError.msg)
//		return
//	}
//	fmt.Println("caiu fora")
//}

// Quarto exemplo de ERRO //

//sem error warpping

//func foo() error {
//	err := bar()
//	if err != nil {
//		return errors.New("deu erro em foo" + err.Error())
//	}
//	return nil
//}

// com error warpping

//func foo() error {
//	err := bar()
//	if err != nil {
//		return fmt.Errorf("deu erro em foo: %w", err)
//	}
//	return nil
//}
//
//var ErrQualquer = errors.New("error")
//
//func bar() error { return ErrQualquer }
//
//func main() {
//	err := foo()
//	if err != nil && errors.Is(err, ErrQualquer) {
//		fmt.Println("deu erro:", err)
//		return
//	}
//}

// Quinto exemplo de ERRO - Agrupando erros//

var (
	ErrQualquer  = errors.New("error")
	ErrQualquer2 = errors.New("error2")
)

func a() error { return ErrQualquer }
func b() error { return ErrQualquer2 }

func foo() error {
	var errorResult error
	if err := a(); err != nil {
		errorResult = errors.Join(errorResult, err)
	}

	if err := b(); err != nil {
		errorResult = errors.Join(errorResult, err)
	}
	return errorResult
}

func main() {
	err := foo()
	fmt.Println(err)
	fmt.Println(errors.Is(err, ErrQualquer))
	fmt.Println(errors.Is(err, ErrQualquer2))
}
