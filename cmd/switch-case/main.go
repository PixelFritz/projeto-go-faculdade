package main

import (
	"fmt"
)

func main() {
	var opcao int

	fmt.Println("--------- Menu ---------")
	fmt.Println("Escolha uma opção da lista:")
	fmt.Println("1 - Imprima seu nome")
	fmt.Println("2 - Imprima seu nome do meio")
	fmt.Println("3 - Imprima seu último nome")
	fmt.Println("4 - Imprima seu apelido")
	fmt.Println("5 - Imprima o nome do seu pet")

	fmt.Print("\nOpção desejada: ")
	fmt.Scanln(&opcao)

	var texto string

	switch opcao {
	case 1:
		fmt.Print("Escreva o seu nome: ")
		fmt.Scanln(&texto)
		fmt.Println("Seu nome é: " + texto)
	case 2:
		fmt.Print("Escreva o seu nome do meio: ")
		fmt.Scanln(&texto)
		fmt.Println("Seu nome do meio é: " + texto)
	case 3:
		fmt.Print("Escreva o seu último nome: ")
		fmt.Scanln(&texto)
		fmt.Println("Seu último nome é: " + texto)
	case 4:
		fmt.Print("Escreva o seu apelido: ")
		fmt.Scanln(&texto)
		fmt.Println("Seu apelido é: " + texto)
	case 5:
		fmt.Print("Escreva o nome do pet: ")
		fmt.Scanln(&texto)
		fmt.Println("O nome do seu pet é: " + texto)
	default:
		fmt.Println("Opção inexiste. Saindo...")
	}
}