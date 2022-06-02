package desafio

import (
	"fmt"
	"testing"
)

func TestDesafio(t *testing.T) {

	nome := []string{
		"c",
		"a",
		"r",
		"l",
		"o",
		"s",
		" ",
		"f",
		"e",
		"l",
		"i",
		"p",
		"e",
	}

	fmt.Println(nome)
	nos := CriarNos(nome)

	primeiro := nos[0]
	primeiro.Add(nos[1:]...)

	esperado := "carlos felipe"
	recebido := Imprimir(primeiro)

	if esperado != recebido {
		t.Errorf("Esperado %s, recebido %s", esperado, recebido)
	}
}
