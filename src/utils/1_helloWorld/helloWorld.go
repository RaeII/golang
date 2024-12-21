package helloWorld

import "fmt"

func DaleHelloWord() {

	fmt.Printf("\n\n------------HELLOWORLD-----------\n\n");

	/*
		func Println(a ...any) (n int, erro error)
		"..." indica que é uma função variádica(variadic), pode receber varios argumentos
		O uso "_" indica que não será usado um dado de retorno, como no "println", não está sendo usado o retorno de erro
	*/
	numbersBytes, _ := fmt.Println("Hello world!");
	fmt.Println(numbersBytes)
}

