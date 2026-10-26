package main

import (
	"fmt"
	"os"
)

func main() {

	// --- Echo1: implementação do comando echo do Unix, com laço for clássico ---
	// s acumula o texto final; sep é o separador entre os argumentos.
	// `var` sem valor inicial dá a ambas o valor zero de string: "".
	var s, sep string

	// os.Args[0] é o nome do próprio programa, então começamos em 1.
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		// Só depois do primeiro argumento o separador passa a ser um espaço,
		// evitando um espaço sobrando no início.
		sep = " "
	}
	fmt.Println(s)
}
