package arrray

import "fmt"

// iota é um incrementador de contantes de números inteiros

var x [5]int

func Array() {

	fmt.Printf("\n\n------------ ARRAY -----------\n\n");

	x[4] = 17
	x[0] = 6

	fmt.Println(x)
	fmt.Println("Tamanho array", len(x))
	
}

