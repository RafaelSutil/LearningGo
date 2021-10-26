package main

import (
	"fmt"
)

func main() {
	slice := []int{8, 2, 6, 4, 5, 9, 1, 7, 3}
	fmt.Println("Vetor antes")
	fmt.Println(slice)
	fmt.Println("\nVetor depois")
	fmt.Println(quicksort(slice))
}

func quicksort(slice []int) []int {
	//fmt.Println(slice)
	//time.Sleep(time.Second * 3)

	if len(slice) <= 1 {
		return slice
	}

	index := len(slice)
	pivo := slice[index/2]
	//fmt.Println(index)
	//fmt.Println(pivo)
	maiores, menores := particionar(slice, pivo)
	//fmt.Println(maiores)
	//fmt.Println(menores)
	//fmt.Printf("\n\n")

	return append(append(quicksort(menores), pivo), quicksort(maiores)...)
}

func particionar(slice []int, pivo int) ([]int, []int) {
	maiores := []int{}
	menores := []int{}

	for _, v := range slice {
		if v > pivo {
			maiores = append(maiores, v)
		}
		if v < pivo {
			menores = append(menores, v)
		}
	}
	return maiores, menores
}
