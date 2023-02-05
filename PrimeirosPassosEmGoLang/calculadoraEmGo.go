package main // pacote principal

import "fmt" //importar pacotes

//variáveis
var opcao int
var resultado float32
var valor1 float32
var valor2 float32

//calculadora
func main() {
	fmt.Println("Calculadora em GO Lang\n")
	//coletar a opção do usuário
	fmt.Println("Digite um número para realizar uma operação matemática")
	fmt.Println("1 - Soma")
	fmt.Println("2 - Subtração")
	fmt.Println("3 - Divisão")
	fmt.Println("4 - Multiplicação")
	fmt.Scan(&opcao)

	//coletar os valores para realizar a operação
	fmt.Println("Digite o primeiro valor")
	fmt.Scan(&valor1)

	fmt.Println("Digite o segundo valor")
	fmt.Scan(&valor2)

	//de acordo com a opção artimética escolhida, chamar a função correspondente e enviar como parâmetro os 2 valores coletados
	switch opcao {
	case 1:
		resultado = Somar(valor1, valor2)
	case 2:
		resultado = Subtrair(valor1, valor2)
	case 3:
		resultado = Dividir(valor1, valor2)
	case 4:
		resultado = Multiplicar(valor1, valor2)
	default:
		fmt.Println("erro") //se a opção não for válida
	}

	//exibe o resultado retornado
	fmt.Printf("O resultado de seu cálculo foi %.2f", resultado)
}

//função de soma
func Somar(valor1 float32, valor2 float32) float32 {
	return valor1 + valor2
}

//função de subtração
func Subtrair(valor1 float32, valor2 float32) float32 {
	return valor1 - valor2
}

//função de divisão
func Dividir(valor1 float32, valor2 float32) float32 {
	return valor1 / valor2
}

//função de mutiplicação
func Multiplicar(valor1 float32, valor2 float32) float32 {
	return valor1 * valor2
}

//para rodar, digitar no terminal:  go run calculadoraEmGo.go

/*
string = %s
integer = %d
float = %f
exibir o tipo da variável = %T (T maiúsculo)
*/
