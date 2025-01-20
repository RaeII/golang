package interfaces

import "fmt"

type piranha struct{
	name string
	lastName string
	level string
	age int
}

type dentista struct {
	piranha piranha
	dentesArrancados int
	salario float64
}

type arquiteto struct {
	piranha piranha
	tipoConstrução string
	nivelLoucura string
}

func (x dentista) olaBomdia(){
	fmt.Println("Ola bom dia meu nome é", x.piranha.name, "e já arranquei",x.dentesArrancados,"dentes")
}

func (x arquiteto) olaBomdia(){
	fmt.Println("Ola bom dia meu nome é", x.piranha.name)
}



func Interfaces() {

	mrDente := dentista{
		piranha: piranha{
			"Vih",
			"Prado",
			"Vale do Guaiba",
			24,
		},
		dentesArrancados:32,
		salario: 1526.30,
	}

	mrObra := arquiteto{
		piranha: piranha{
			"Vitória",
			"Minossi",
			"Vale do itajaí",
			25,
		},
		tipoConstrução: "Fundação",
		nivelLoucura: "Um pouco fora",
	}

	mrDente.olaBomdia()
	mrObra.olaBomdia()

	fmt.Printf("\n\n------------ INTERFACE -----------\n\n");



	
}

