package main

import (
	"fmt"
)

// Mapas

// o valor do indice 0 de um mapa é nulo, assim como o valor 0 de um slice
// mapas em go são desordenados

//func main() {
//	var m map[string]string
//	fmt.Println(m == nil)
//}

//func main() {
//	m := make(map[string]string)
//	fmt.Println(m == nil)
//}

//func main() {
//	m := map[string]string{
//		"Pedro":   "Pessoa",
//		"Joaquim": "Pedro",
//	}
//	fmt.Println(m == nil)
//	fmt.Println(m)
//}

// Mapa de Slice

//func main() {
//	m := map[string][]int{
//		"Pedro": {1, 2, 3},
//	}
//	fmt.Println(m)
//}

// obtendo um elemento de um mapa
//func main() {
//	m := make(map[string]string)
//	m["Pedro"] = "Pessoa"
//	valor, ok := m["Pedro"] // ok diz se a chave existe
//	fmt.Println(valor, ok)
//	delete(m, "Pedro")
//	valor, ok = m["Pedro"]
//	fmt.Println(valor, ok)
//}

//retirando uma chave Nan de um mapa

//func main() {
//	f := math.NaN()
//	f2 := math.NaN()
//	m := map[float64]string{
//		f:  "Pedro",
//		f2: "[Pessoa]",
//	}
//	fmt.Println(m)
//	valor, ok := m[f]
//	fmt.Println(valor, ok)
//	delete(m, f)
//	fmt.Println(m)
//	clear(m)
//	fmt.Println(m)
//}

// iterando em um mapa

func main() {
	m := map[string]string{
		"Pedro":   "Pessoa",
		"Joaquim": "Pedro",
	}
	//for k, v := range m {
	//	fmt.Println(k, v)
	//}

	for k := range m {
		if k == "Pedro" {
			delete(m, k)
		}
	}
	fmt.Println(m)
}
