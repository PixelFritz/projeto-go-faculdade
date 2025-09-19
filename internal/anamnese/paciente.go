package anamnese

import (
	"fmt"
	"math"
)

type Paciente struct {
	Nome string
	Peso float64
	Altura float64
}

func CalcularIMC(peso float64, altura float64) float64 {
	resultadoIMC := peso / (math.Pow(altura, 2))
	return resultadoIMC
}

func VerificarIMC(resultadoIMC float64, nome string) string {
	var resposta string

	switch {
	case resultadoIMC < 18.5:
		resposta = fmt.Sprintf("%v está abaixo do peso normal", nome)
	case resultadoIMC >= 18.5 && resultadoIMC <= 24.9:
		resposta = fmt.Sprintf("%v está no peso normal", nome)
	case resultadoIMC >= 25 && resultadoIMC <= 29.9:
		resposta = fmt.Sprintf("%v está com excesso de peso", nome)
	case resultadoIMC >= 30 && resultadoIMC <= 34.9:
		resposta = fmt.Sprintf("%v está com obesidade tipo I", nome)
	case resultadoIMC >= 35 && resultadoIMC <= 39.9:
		resposta = fmt.Sprintf("%v está com obesidade tipo II", nome)
	case resultadoIMC >= 40:
		resposta = fmt.Sprintf("%v está com obesidade tipo III", nome)
	default:
		resposta = fmt.Sprintf("Erro: IMC inválido para %v", nome)
	}

	return resposta
}
