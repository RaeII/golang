package types

import "fmt"

type bodyRequest int

var body bodyRequest = 10

func Types() {

	newBody := 32

	fmt.Printf("\n\n------------ TYPES -----------\n\n");

	fmt.Printf("tipo do body : %T\n",body)
	fmt.Printf("tipo do new body : %T\n",newBody)
	
	//Valor só pode ser atribuido ao mesmo tipo
	//body = newBody

	//precisa fazer conversão(conversion) dos tipos 
	body = bodyRequest(newBody)
	
}

