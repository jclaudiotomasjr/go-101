package main

import "fmt"

func main() {

	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}
	fmt.Println("*******************************************")
	nomes := []string{"Vitor", "Yan", "Stefania", "Jujuba"}
	for k, nome := range nomes {
		fmt.Println(k, nome)
	}

	fmt.Println("*******************************************")
	var count int
	for count < len(nomes) {
		fmt.Println(nomes[count])
		count++
	}

}
