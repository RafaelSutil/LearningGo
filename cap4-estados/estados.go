package main

import "fmt"

type Estado struct {
	nome      string
	populacao int
	capital   string
}

func main() {
	estados := make(map[string]Estado, 6)
	// Populando
	estados["GO"] = Estado{"Goiás", 6500000, "Goiânia"}
	estados["PB"] = Estado{"Paraíba", 4000000, "João Pessoa"}
	estados["RN"] = Estado{"Rio Grande do Norte", 3500000, "Natal"}
	estados["AM"] = Estado{"Amazonas", 3800000, "Manaus"}
	estados["SE"] = Estado{"Sergipe", 2200000, "Aracaju"}
	estados["PR"] = Estado{"Paraná", 11000000, "Curitiba"}

	fmt.Println(estados)

	// Lookup
	sergipe := estados["SE"]
	fmt.Println(sergipe)

	pernambuco, encontrado := estados["PE"]
	if encontrado {
		fmt.Println(pernambuco)
	}

	// Delete
	delete(estados, "PR")
	fmt.Println(estados)

	// Iteração
	for sigla, estado := range estados {
		fmt.Printf("%s (%s) possui %d habitantes.\n", estado.nome, sigla, estado.populacao)
	}

}
