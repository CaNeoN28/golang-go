package main

import "testing"

func TestBubbleSort(t *testing.T){
	lista := []int{9,8,7,6,5,4,3,2,1,0}

	resultado := BubbleSort(lista)

	esperado := []int{0,1,2,3,4,5,6,7,8,9}

	for idx := range esperado{
		item1 := resultado[idx]
		item2 := esperado[idx]

		if item1 != item2{
			t.Error("A função não funciona como esperado")
			break
		}
	}
}