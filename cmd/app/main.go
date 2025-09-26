// Arquivo principal do programa (entrypoint)
// Convenção de mercado: colocar em cmd/<nome-app>/main.go
package main

// Importa os pacotes necessários
import (
	"fmt"

	"github.com/seu-usuario/meu-projeto-go/internal/anamnese"
	"github.com/seu-usuario/meu-projeto-go/internal/fibonacci"
	"github.com/seu-usuario/meu-projeto-go/internal/hello"
	"github.com/seu-usuario/meu-projeto-go/internal/saudacao"
	"github.com/seu-usuario/meu-projeto-go/internal/funcao-anonima"
)

// Função principal do programa
func main() {
	fmt.Println("Escolha uma opção:")
	fmt.Println("1 - Hello")
	fmt.Println("2 - Fibonacci")
	fmt.Println("3 - Anamnese")
	fmt.Println("4 - Saudação")
	fmt.Println("5 - Função Anônima")

	var opcao int
	fmt.Print("Digite o número da opção desejada: ")
	fmt.Scanln(&opcao)

	switch opcao {
	case 1:
		fmt.Println("🚀 Meu primeiro projeto em Go com estrutura de mercado!")
		hello.SayHello()
	case 2:
		var entradaFibonacci int
		fmt.Print("\nEscolha um número para calcular o Fibonacci: ")
		fmt.Scanln(&entradaFibonacci)

		if entradaFibonacci < 0 {
			entradaFibonacci = -entradaFibonacci
		}

		result := fibonacci.Fibonacci(entradaFibonacci)
		fmt.Printf("O resultado do Fibonacci(%v) é: %v\n", entradaFibonacci, result)
	case 3:
		var pac anamnese.Paciente

		fmt.Print("\nEscreva o seu nome: ")
		fmt.Scanln(&pac.Nome)

		fmt.Print("Escreva o seu peso (kg): ")
		fmt.Scanln(&pac.Peso)

		fmt.Print("Escreva a sua altura (m): ")
		fmt.Scanln(&pac.Altura)

		resultadoIMC := anamnese.CalcularIMC(pac.Peso, pac.Altura)
		fmt.Printf("\nO resultado do IMC é: %.2f\n", resultadoIMC)

		classificacao := anamnese.VerificarIMC(resultadoIMC, pac.Nome)
		fmt.Println(classificacao)
	case 4:
		var nome string

		fmt.Print("\nEscreva o seu nome: ")
		fmt.Scanln(&nome)

		fmt.Println(saudacao.Saudacao(nome))
	case 5:
		funcaoanonima.PreencheVetor()
	default:
		fmt.Println("Número inválido")
	}
}
