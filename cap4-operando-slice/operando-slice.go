package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7}
	// --- FATIAR --- ///
	fmt.Println("---- FATIAR ----")

	novoSlice := slice[1:4]

	// antes de modificar
	fmt.Println("Antes")
	fmt.Println("Slice: ", slice)
	fmt.Println("novoSlice: ", novoSlice)

	slice[3] = 99

	// apos modificar
	fmt.Println("Depois")
	fmt.Println("Slice: ", slice)
	fmt.Println("novoSlice: ", novoSlice)

	// --- INSERIR --- ///
	fmt.Println("---- INSERIR ----")
	// inserir no final
	slice = append(slice, 8)
	fmt.Println("Slice: ", slice)

	//inserir no inicio
	n := []int{0}
	slice = append(n, slice...)
	fmt.Println("Slice: ", slice)

	// --- REMOVENDO --- ///
	fmt.Println("---- REMOVENDO ----")
	slice = append(slice[:2], slice[4:]...)
	fmt.Println("Slice: ", slice)

	// --- COPIANDO --- ///
	fmt.Println("---- COPIANDO ----")
	s := make([]int, len(slice))
	copy(s, slice)
	for i := range s {
		s[i] += 1
	}
	fmt.Println("s: ", s)
	fmt.Println("Slice: ", slice)
}
