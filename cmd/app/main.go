// Arquivo principal do programa (entrypoint)
// Convenção de mercado: colocar em cmd/<nome-app>/main.go
package main

// Importa os pacotes necessários
import (
	"fmt"
	"os"

	"github.com/seu-usuario/meu-projeto-go/internal/anamnese"
	"github.com/seu-usuario/meu-projeto-go/internal/fibonacci"
	"github.com/seu-usuario/meu-projeto-go/internal/hello"
)

// Função principal do programa
func main() {
	fmt.Println("Escolha uma opção:")
	fmt.Println("1 - Hello")
	fmt.Println("2 - Fibonacci")
	fmt.Println("3 - Anamnese")

	var opcao int
	fmt.Print("Digite o número da opção desejada: ")
	fmt.Fscan(os.Stdin, &opcao)

	switch opcao {
	case 1:
		fmt.Println("🚀 Meu primeiro projeto em Go com estrutura de mercado!")
		hello.SayHello()
	case 2:
		var entradaFibonacci int
		fmt.Print("\n\nEscolha um número para calcular o Fibonacci: ")
		fmt.Fscan(os.Stdin, &entradaFibonacci)
		result := fibonacci.Fibonacci(entradaFibonacci)
		fmt.Printf("O resultado do Fibonacci(%v) é: %v\n", entradaFibonacci, result)
	case 3:
		var pac anamnese.Paciente

		fmt.Print("Escreva o seu nome: ")
		fmt.Fscan(os.Stdin, &pac.Nome)
		fmt.Print("Escreva o seu peso: ")
		fmt.Fscan(os.Stdin, &pac.Peso)
		fmt.Print("Escreva a sua altura: ")
		fmt.Fscan(os.Stdin, &pac.Altura)

		resultadoIMC := anamnese.CalcularIMC(pac.Peso, pac.Altura)
		fmt.Printf("\nO resultado do IMC é: %.2f\n", resultadoIMC)

		classificacao := anamnese.VerificarIMC(resultadoIMC, pac.Nome)
		fmt.Println(classificacao)
	default:
		fmt.Println("Número inválido")
	}
}
