package strucs

import "fmt"

type person struct {
	name string
	lastName string
	age int
	cpf string
}

type piranha struct {
	person person
	nivel string
}

func Struct() {

	fmt.Printf("\n\n------------ Struc -----------\n\n");

	person1 := person{
		name:"Israel",
		lastName: "Zeferino",
		age: 28,
		cpf: "098.079.456.77",
	}

	person2 := person{
		"Vitoria",
		"Prado",
		25,
		"125.465.456.77",
	}

	piranha1 := piranha{
		person2,
		"Vale do Guaíba",
	}

	piranha2 := piranha{
		person{
			"Vitoria",
			"Minossi",
			26,
			"125.465.456.77",
		},
		"Vale do Itajai",
	}

	fmt.Println(person1)
	fmt.Println(person2)
	fmt.Println(piranha1)
	fmt.Println(piranha2)

	piranhaAnonima := struct {
		person person
		nivel string
	}{
		person{
			"Vitoria",
			"Anonima",
			26,
			"125.465.456.77",
		},
		"Vale do anonimato",
	}

	fmt.Println("\n\nPiranha anonima",piranhaAnonima)
}

