package fatia

import "fmt"


func Fatia() {

	fmt.Printf("\n\n------------ FATIA -----------\n\n");

	city := []string{"Penha","Navega","Itaja","Bc","Ita","PortoBe","Floripa","Sombrio"}

	trexo := city[:3]
	fmt.Println("trexo 1", trexo)

	trexo2 := city[3:7]
	fmt.Println("trexo 2", trexo2)

	trexo3 := city[7:8]
	fmt.Println("trexo 3", trexo3)

	trexo4 := city[7:]
	fmt.Println("trexo 4", trexo4)

	fmt.Println("Tamanho do trexo",len(city))

	//DELETAR
	city = append(city[:6],city[7:]...)
	fmt.Println("New trecho", city)

}

