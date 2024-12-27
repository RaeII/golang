package forloop

import "fmt"

// iota é um incrementador de contantes de números inteiros

func Forloop() {

	fmt.Printf("\n\n------------ FOR -----------\n\n");

	for x := 0; x < 20; x++ {
		fmt.Println("x: ",x)
	}

	//Em GO não existe while

	x := 1
	for(x <= 10){
		println("X :=",x)
		x++
	}
	
}

