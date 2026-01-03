package main

import "fmt"

var filmesNoDB = []string{
	"O Poderoso Chefão",
	"Titanic",
	"O Senhor dos Anéis: O Retorno do Rei",
	"Matrix",
	"Forrest Gump: O Contador de Histórias",
	"O Rei Leão",
	"Harry Potter e a Pedra Filosofal",
	"Gladiador",
	"O Sexto Sentido",
	"O Curioso Caso de Benjamim Button",
	"Pulp Fiction: Tempo de Violência",
	"Esqueceram de Mim",
	"O Exterminador do Futuro 2: O Julgamento Final",
	"O Fabuloso Destino de Amélie Poulain",
	"O Labirinto do Fauno",
}

func secondmain() {
	resultsFromApi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	//var filmes []string

	//slice pre alocado

	filmes := make([]string, 0, 10)

	for _, id := range resultsFromApi {
		filme := filmesNoDB[id]
		fmt.Println(len(filmes), cap(filmes))
		filmes = append(filmes, filme)
		fmt.Println(len(filmes), cap(filmes))
	}
	fmt.Println(filmes)

	//slice 2D e 3D
	//matrix := [][]int{}
	//matrix3D := [][][]int{}
}
