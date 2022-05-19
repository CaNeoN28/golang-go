package main

import "fmt"

func main() {
	var raio float32
	fmt.Scan(&raio)
	
	area := calcularArea(raio)
	fmt.Printf("A=%.4f\n", area)
}