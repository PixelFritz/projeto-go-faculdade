// Arquivo principal do programa (entrypoint)
// Convenção de mercado: colocar em cmd/<nome-app>/main.go
package main

// Importa os pacotes necessários
import (
	"fmt"
	"os"

	"github.com/seu-usuario/meu-projeto-go/internal/fibonacci"
	"github.com/seu-usuario/meu-projeto-go/internal/hello"
)

// Função principal do programa
func main() {
    fmt.Println("🚀 Meu primeiro projeto em Go com estrutura de mercado!")
    hello.SayHello()

    var entradaFibonacci int
    fmt.Print("\n\nEscolha um número para calcular o Fibonacci: ")
	fmt.Fscan(os.Stdin, &entradaFibonacci)
    result := fibonacci.Fibonacci(entradaFibonacci)
	fmt.Printf("O resultado do Fibonacci(%v) é: %v\n", entradaFibonacci, result)
}
