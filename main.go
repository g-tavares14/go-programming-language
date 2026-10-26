package main

import (
	"fmt"
	"os"
	"strings"
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

	// --- Echo2: o mesmo com range ---
	// Reaproveitamos s e sep com `=` (e não `:=`, que daria erro: sem variáveis novas).
	s, sep = "", ""
	// range devolve (índice, valor) a cada volta. O índice não é usado aqui;
	// o identificador em branco `_` o descarta, já que Go não permite
	// variáveis locais declaradas e não usadas.
	// os.Args[1:] pula o índice 0 (nome do programa).
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)

	// --- Echo3: strings.Join ---
	// Junta os elementos do slice com o separador dado. É mais eficiente que
	// `s += ...` em laço: como strings são imutáveis, cada concatenação cria
	// uma string nova e copia a anterior.
	fmt.Print(strings.Join(os.Args[1:], " "))
}
