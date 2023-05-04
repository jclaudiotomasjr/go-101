package main

import (
	"fmt"
)

func main() {

	//criando uma slice de nomes atribuindo 4 nomes
	nomes := []string{"Jujuba", "Jack", "Johnny", "Curupaco"}
	fmt.Println(nomes, len(nomes), cap(nomes))
	//adicionando mais nomes para o slice nomes com append
	nomes = append(nomes, "Costela")
	fmt.Println(nomes, len(nomes), cap(nomes))
	nomes = append(nomes, "Rocky")
	fmt.Println(nomes, len(nomes), cap(nomes))
	//Criando um slice sem atribuição
	idades := make([]uint8, 0)
	fmt.Println(idades, len(idades), cap(idades))

	fmt.Println("******************************************")
	//Criando um map atribuindo os valores
	idades2 := make(map[string]uint8)
	idades2["Jujuba"] = 3
	idades2["Jack"] = 4
	idades2["Johnny"] = 3
	idades2["Curupaco"] = 1
	//Observação um map não é indexado
	fmt.Println(idades2)
	//validar se um indice existe em um map
	val, ok := idades2["Junior"]
	fmt.Println(val, ok)

}
