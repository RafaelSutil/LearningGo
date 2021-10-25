package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	var i int
	n := 0
	for {
		n++
		i = rand.Intn(4200)
		fmt.Println((i))

		if i%42 == 0 {
			break
		}
	}
	fmt.Printf("Saída após %d iterações.\n", n)

	// break para um loop externo
externo:
	for {
		for i = 0; i < 10; i++ {
			if i == 5 {
				break externo
			}
		}
	}
	fmt.Println(i)
}
