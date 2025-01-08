package subjacente

import "fmt"


func SliceMult() {

	fmt.Printf("\n\n------------ SLICE MULT -----------\n\n");

	x := [][][]int{
		{
			{1, 2, 3},    // Este é um []int
			{7, 8, 9},    // Este é outro []int
		},
		{
			{4, 5, 6},
		},
		{
			{1, 2},
		},
	}

	fmt.Println(x)

}

