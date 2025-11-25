package main

import "fmt"

func main() {
	fmt.Println("insira o valor do primeiro numero:")
	var num1 float64
	fmt.Scanln(&num1)

	fmt.Println("insira o valor do segundo numero:")
	var num2 float64
	fmt.Scanln(&num2)

	fmt.Println("Informe qual operação deseja ser feita: \n1 - Soma \n2 - Subtração \n3 - Multiplicação \n4 - Divisão")
	var operacao int
	fmt.Scanln(&operacao)
	switch operacao {
	case 1:
		resultado := num1 + num2
		fmt.Println("Resultado:", resultado)
	case 2:
		resultado := num1 - num2
		fmt.Println("Resultado:", resultado)
	case 3:
		resultado := num1 * num2
		fmt.Println("Resultado:", resultado)
	case 4:
		if num2 != 0 {
			resultado := num1 / num2
			fmt.Println("Resultado:", resultado)
		} else {
			fmt.Println("Erro: Divisão por zero não é permitida.")
		}
	default:
		fmt.Println("Operação inválida.")
	}
}
