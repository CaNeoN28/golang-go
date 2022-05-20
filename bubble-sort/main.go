package main

import (
	"fmt"
	"math/rand"
	"time"
)

func BubbleSort(array []int) []int {
	copy := array

	for i := 0; i < 1; {
		ocurrences := 0
		for j := 0; j < len(copy); j++ {
			if j != len(copy)-1 && copy[j] > copy[j+1] {
				ocurrences += 1
				aux := copy[j]
				copy[j] = copy[j+1]
				copy[j+1] = aux
			}
		}

		if ocurrences == 0 {
			i += 1
		}
	}

	return copy
}

func main() {
	var array []int

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New((s1))

	for i := 0; i < 10; i++ {
		array = append(array, r1.Intn(100))
	}

	fmt.Printf("Array não ordenado: %d\n", array)
	fmt.Printf("Array ordenado: %d\n", BubbleSort(array))
}
