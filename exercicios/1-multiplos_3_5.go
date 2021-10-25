package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Entrada incorreta!")
		os.Exit(1)
	}

	number, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Entrada deve ser um inteiro!")
		os.Exit(1)
	}

	sum := 0
	multiples := make([]int, 0)
	for i := 0; i < number; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
			multiples = append(multiples, i)
		}
	}

	fmt.Printf("Soma = %d\nValores = %d\n", sum, multiples)
}
