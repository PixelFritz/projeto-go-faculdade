package funcaoanonima

import "fmt"

// Funções anônimas são funções não nomeadas, geralmente utilizadas para coisas rápidas, como callbacks e expressões simples. São utilizadas geralmente dentro do mesmo arquivo e as variáveis podem receber outros valores.

var numeros = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

var dobrar = func(x int) int {
	return x * 2
}

var dobrados = []int{}

func PreencheVetor() {
	for _, n := range numeros {
		dobrados = append(dobrados, dobrar(n))
	}

	fmt.Println("\nOriginais:", numeros)
	fmt.Println("Dobrados:", dobrados)
}
