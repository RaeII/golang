package idiota

import "fmt"

// iota é um incrementador de contantes de números inteiros

func Iota() {

	fmt.Printf("\n\n------------ IOTA -----------\n\n");

	const (
		one = iota + 1
		_
		three = 3
		four = 4
		five = 5
		six = 6
		seven = 7
		eight = 8
		_
		ten = 10
	)

	println(one,three,four,five,six,seven,eight,ten)
	
}

