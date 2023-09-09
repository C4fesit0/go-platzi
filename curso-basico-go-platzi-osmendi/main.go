package main

import (
	"fmt"
)

func normalFunction(message string) {
	fmt.Println(message)
}

func tripleArgument(a, b int, c string) {
	fmt.Println(a, b, c)
}

func returnValue(a int) int {
	return a * 2
}

func dobleReturn(a int) (c, d int) {
	return a, a * 2
}

func main() {

	/*
		normalFunction("Hola Mundo")
		tripleArgument(1, 2, "Hola")

		value := returnValue(2)
		fmt.Println("Value:", value)

		value1, _ := dobleReturn(2)
		fmt.Println("Value1 y Value2:", value1)
		   const pi float64 = 3.14

		   const pi2 = 3.1415

		   fmt.Println("pi:", pi)
		   fmt.Println("pi2:", pi2)

		   //Declaracion de variables enteras

		   base := 12
		   var altura int = 14
		   var area int

		   fmt.Println(base, altura, area)

		   //Zero values
		   var a int
		   var b float64
		   var c string
		   var d bool

		   fmt.Println(a, b, c, d)

		   //Area de un cuadrado

		   const baseCuadrado = 10
		   areaCuadrado := baseCuadrado * baseCuadrado
		   fmt.Println("Area Cuadrado=", areaCuadrado)

		   //Operaciones Aritmeticas

		   x := 10
		   y := 50

		   //Suma
		   result := x + y
		   fmt.Println("Suma:", result)

		   //Resta
		   result = x - y
		   fmt.Println("Resta:", result)

		   //Multiplicacion
		   result = x * y
		   fmt.Println("Multiplicacion:", result)

		   //Division
		   result = x / y
		   fmt.Println("Division:", result)

		   //Modulo
		   result = x % y
		   fmt.Println("Modulo:", result)

		   //Incrementable
		   x++
		   fmt.Println("Incremental:", x)

		   //Decrementable
		   x--
		   fmt.Println("Decremental:", x)

		   /*
		   *Datos Primitivos:
		   *int64,float64, uint64
		   *string --> Se declaran si o si con " " no con ' '
		   *bool
		   *Datos Complejos:
		   *Complex64 -->Real e imaginario float32
		   *Complex128 --> Real e imaginario float64
		   *Ej: c:=10+8i

		   helloMessage := "Hello"
		   worldMessage := "World"

		   //Println
		   fmt.Println(helloMessage, worldMessage)
		   fmt.Println(helloMessage, worldMessage)

		   //Printf
		   nombre := "Platzi"
		   cursos := 500
		   fmt.Printf("%s tiene mas de %d cursos \n", nombre, cursos)
		   //Si no sabemos el tipo de dato, podemos usar %v
		   //La buena practica es agregar el tipo de dato
		   fmt.Printf("%v tiene mas de %v cursos \n", nombre, cursos)

		   //Sprintf
		   message := fmt.Sprintf("%v tiene mas de %v cursos \n", nombre, cursos)
		   fmt.Print(message)

		   // Tipo de dato
		   fmt.Printf("helloMessage: %T \n", helloMessage)
		   fmt.Printf("cursos: %T", cursos)
	*/
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
