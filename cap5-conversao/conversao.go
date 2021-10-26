package main

import "fmt"

type ListaDeCompras []string

func imprimirSlice(slice []string) {
	fmt.Println("Slice:", slice)
}

func imprimirLista(lista ListaDeCompras) {
	fmt.Println("Lista de compras:", lista)
}

func main() {
	lista := ListaDeCompras{"Arroz", "Tomate", "Sorvete"}
	slice := []string{"Arroz", "Tomate", "Sorvete"}

	imprimirSlice([]string(lista))
	imprimirLista(ListaDeCompras(slice))
}
