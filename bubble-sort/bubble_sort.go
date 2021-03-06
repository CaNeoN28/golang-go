package main

import (
	"fmt"
	"math/rand"
	"time"
)

func BubbleSort(array []int) []int {
	copy := array

	ocurrences := 0
	for j := 0; j < len(copy) - 1; j++ {
		if copy[j] > copy[j+1] {
			ocurrences += 1
			copy[j], copy[j + 1] = copy[j+1], copy[j]
		}
	}

	if ocurrences != 0 {
		copy = BubbleSort(copy)
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
