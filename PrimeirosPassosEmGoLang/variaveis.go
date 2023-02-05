// pacote principal
package main

import ( //importar pacotes
	"fmt" //pacote format com a função de printar na tela (Println e Printf)
)

//declarações de variáveis
//a linguagem GO é fortemente tipada, não sendo possível modificar o tipo após a atribuição
//tipos de variáveis em GO: int, float (e suas variações), string, bool

//toda variável declarada deve ser obrigatoriamente utilizada

//modelo de declaraçao e atribuição
//var nomeDaVariavel tipo = valor

//curiosidade do atribuidor curto de variável := marmota ~ Gopher - simbolo do GO
//o := possui tipagem automática

var nome string //declaração da variável nome
//como a variavel nome está fora de uma função, ela possui package level scope, é como se fosse uma "variável global"
//sendo de abrangencia em qualquer local dentro do pacote

// testeComMarmota := "valor" - não é possível utilizar a marmota fora de um codeblock (função)

func main() {
	nome = "Lidiaa"   //atribuição de valor a variável nome
	nome2 := "github" //criação e atribuição da variável nome2, que pelo conteúdo recebido irá presumir o tipo de dados
	//a variavel nome2 foi criada dentro da função main, ou seja, apenas funcionará dentro desse bloco de função - scope

	fmt.Println(nome)  //print Lidiaa
	fmt.Println(nome2) //print github

	nome = nome2      //alteração do valor da variavel nome que recebe o conteúdo de nome2, ou seja, a palavra "github"
	fmt.Println(nome) //print github também

	//declarações e atribuições de valores a variáveis
	valor1 := 10      //int - número inteiro (sem vírgukas)
	valor2 := "World" //string - palavras ou letras
	valor3 := 3.144   //float64 - número com vírgula
	valor4 := false   //bool - booleano (verdadeiro ou falso)
	valor5 := `usando o 
	acento
	grave` //string - palavars ou letras

	//exibindo as variáveis anteriormente criadas
	fmt.Println("\n\n== Valores: ==")

	//ao utilizar %v, é exibido na tela o Valor atribuido a elas
	fmt.Printf("%v \n", valor1)
	fmt.Printf("%v \n", valor2)
	fmt.Printf("%v \n", valor3)
	fmt.Printf("%v \n", valor4)
	fmt.Printf("%v \n", valor5)

	fmt.Println("\n\n== Tipos: ==")

	//ao utilizar %T, é exibido na tela o Tipo atribuido a elas
	fmt.Printf("%T \n", valor1)
	fmt.Printf("%T \n", valor2)
	fmt.Printf("%T \n", valor3)
	fmt.Printf("%T \n", valor4)
	fmt.Printf("%T \n", valor5)

}

//para rodar, digitar no terminal go run variaveis.go
