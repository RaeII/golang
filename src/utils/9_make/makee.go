package makee

import "fmt"


func Make() {

	/**
	Sempre que é usando o "append" é criado um novo array e decartado o antigo, isso gera custo de maquina.
	Usando "make" é passado um minimo de tamanho e um maximo, assim evitando de criar sempre um novo array.
	Se o tamanho maximo for usado, o valor maximo é dobrado (Não é garantido que vai dobrar, depende das versões do GO)
	*/

	fmt.Printf("\n\n------------ MAKEE -----------\n\n");

	x := make([]int,5,10)
	y := []int{1}

	fmt.Println(cap(y))

	fmt.Println("X ->",x)
	fmt.Println("len ->",len(x))
	fmt.Println("cap ->",cap(x))
	fmt.Println("=============")

	//Como não foi setado valor na posição "[4]", na declaração com "make" já tem valor "0", então no "append" o novo valor será setado depois do "0" na posição "4"
	x[0],x[1],x[2],x[3] = 0,1,2,3

	fmt.Println("X ->",x)
	fmt.Println("len ->",len(x))
	fmt.Println("cap ->",cap(x))
	fmt.Println("=============")


	x = append(x, 5)
	// x = [0 1 2 3 0 4]


	fmt.Println("X ->",x)
	fmt.Println("len ->",len(x))
	fmt.Println("cap ->",cap(x))
	fmt.Println("=============")

	x = append(x, 6)
	x = append(x, 7)
	x = append(x, 8)
	x = append(x, 9)
	x = append(x, 10)

	fmt.Println("X ->",x)
	fmt.Println("len ->",len(x))
	fmt.Println("cap ->",cap(x))
	fmt.Println("=============")
}

