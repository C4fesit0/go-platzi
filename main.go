package main

import (
	"fmt"
)

func main() {
	//For
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	fmt.Print("/n")

	//For while
	counter := 0
	for counter < 10 {
		fmt.Println(counter)
		counter++
	}

	//For forever
	/*counterForever := 0

	 for {
		fmt.Println(counterForever)
		counterForever++
	} */

	//If
	valor1 := 2
	if valor1 == 1 {
		fmt.Println("Es 1")
	} else {
		fmt.Println("No es 1")
	}

	//Array
	var array [4]int
	array[0] = 1
	array[1] = 2
	fmt.Println(array, len(array), cap(array))
	//len -> muesta la longitud del array
	//cap -> muesta cuantos elementos puede guardar un array

}
