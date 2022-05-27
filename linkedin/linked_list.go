package linkedin

import "fmt"

type No struct {
	Valor   int
	Próximo *No
}

func (n *No) Add(prox *No) {
	n.Próximo = prox
}

func Imprimir(n *No) {
	if n != nil {
		fmt.Printf("%d", n.Valor)
	}
}

func main() {
	fmt.Printf("Testando")
}
