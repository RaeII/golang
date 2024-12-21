package variables

import "fmt"

// "x" pode ser acessado nivel pacote, o uso da marmota não pode ser usado fora de escopo
var x = "dale"

func Variables() {

	fmt.Printf("\n\n------------VARIABLES-----------\n\n");
	
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

