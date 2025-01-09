package maps

import "fmt"


func Map() {

	fmt.Printf("\n\n------------ Maps -----------\n\n");

	dogsAge := map[string]int{
		"Connor":6,
		"Pequeno":7,
	}

	fmt.Println(dogsAge)

	dogsAge["Toquinho"] = 0

	fmt.Println(dogsAge)

	//"Ok" retorna se a chave existe nop map
	toquinhoAge, okToquinhoAge := dogsAge["Toquinho"]
	xexeAge, okXexeAge := dogsAge["Xexe"]

	println("Toquinho ->", toquinhoAge,okToquinhoAge)
	println("Xexe ->",xexeAge,okXexeAge)
	print("\n\n")

	//RANGE
	for key,value := range dogsAge {
		fmt.Println("dog ->",key,"idade ->",value)
	}

	print("\n\n")

	//DELETAR
	delete(dogsAge,"Toquinho")

	for key,value := range dogsAge {
		fmt.Println("dog ->",key,"idade ->",value)
	}
}

