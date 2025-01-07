package slicee

import "fmt"

// iota é um incrementador de contantes de números inteiros

var slice = []int{14,45,89,4,6}

func Slice() {

	fmt.Printf("\n\n------------ SLICE -----------\n\n");

	slice[4] = 17
	slice[0] = 6

	fmt.Println(slice)
	fmt.Println("Tamanho slice", len(slice))
	
	slice2 := append(slice,10)
	
	fmt.Println("slice ->", slice)
	fmt.Println("slice 2 ->", slice2)

	fruit := []string{"banana","maça","limão","pera"}

	for index,value := range fruit {

		fmt.Println("index:",index,"value",value)
	}
}

