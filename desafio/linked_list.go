package desafio

import "fmt"

func CriarNos(letras []string) []*No {
	nos := []*No{}

	for _, letra := range letras {
		no := &No{Value: letra}
		nos = append(nos, no)
	}

	return nos
}

type No struct {
	Value   string
	Proximo *No
}

func (n *No) Add(outros ...*No) {
	if len(outros) == 0 {
		return
	}

	primeiro := outros[0]
	n.Proximo = primeiro
	primeiro.Add(outros[1:]...)
}

func Imprimir(n *No) string {
	completo := ""

	for n != nil {
		fmt.Printf("%s ->", n.Value)
		completo += n.Value
		n = n.Proximo
	}

	fmt.Println()

	return completo
}
