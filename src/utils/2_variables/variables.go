package variables

import "fmt"

// "x" pode ser acessado nivel pacote, o uso da marmota não pode ser usado fora de escopo
var x = "dale"

//A definição do tipo de uma constante é feito na hora do uso, no exemplo pode ser um INT ou um FLOAT
const constant = 10
var count float64

//posso declarar varias constantes de uma vez só:
const (
	a = 10
	d = "dale"
	f = 10.10
)

/*

atribuição só pode ser executada dentro de um code block

var y int
y = 10

*/

func Variables() {
	fmt.Printf("\n\n------------VARIABLES-----------\n\n");

	fmt.Println(a,d,f)

	count = constant

	fmt.Printf("constant: %v, %T\n",constant,constant,);
	fmt.Printf("count: %v, %T\n",count,count,);
	
	// ":=" declaração
	y := 1022
	fmt.Printf("x: %v, %T\n",x,x,);
	fmt.Printf("y: %v, %T\n\n",y,y);

	// atribuição, pois "x" já foi declarado
	x = "dalee"
	fmt.Printf("Novo valor de x: %v, %T\n",x,x,);

	//o uso ":=" para uma atribuição pode ser feito se tiver apenas uma declaraão
	x,z := "opa",17
	fmt.Printf("y: %v, %T\n",z,z);
	fmt.Printf("y: %v, %T",x,x);

}

