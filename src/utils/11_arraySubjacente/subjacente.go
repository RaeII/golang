package slicemult

import "fmt"


func Subjacente() {

	fmt.Printf("\n\n------------ Subjacente -----------\n\n");

	slicex := []int{1,2,3,4,5}

	fmt.Println("slicex",slicex)

	/*
		Tem que tomar cuidado ao usar o append, pois ele pode mudar o slice anterior para criar um novo
		Ele muda o anterio pois o Go reaproveira o slice anterior para criar o novo
	*/

	slicey := append(slicex[:2],slicex[4:5]...)

	fmt.Println("slicex",slicex)
	fmt.Println("slicey",slicey)


}

