package main //pacote principal

//importar pacotes
import "fmt" //pacote format com a função de printar na tela (Printf)

func main() {
	resultado := Soma(5, 1) //declaração e atribuição da variável resultado, que recebe o return da função Soma
	fmt.Printf("A soma dos valores foi %v", resultado)

}

func Soma(valor1 int, valor2 int) int { //recebe 2 valores inteiros e retorna um valor int também
	return valor1 + valor2 //retorna a soma dos valores recebidos
}

//para rodar, digitar no terminal go run primeiraFuncao.go
