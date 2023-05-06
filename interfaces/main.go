package main

import "fmt"

type figuras2d interface {
	area() float64
}

type cuadrado struct {
	base float64
}

type rectangulo struct {
	base   float64
	altura float64
}

func (c cuadrado) area() float64 {
	return c.base * c.base
}

func calculate(f figuras2d) {
	fmt.Println("Area:", f.area())
}

func (r rectangulo) area() float64 {
	return r.base * r.altura
}

func main() {
	myCuadrado := cuadrado{base: 2}
	myRectangulo := rectangulo{base: 2, altura: 4}
	calculate(myCuadrado)
	calculate(myRectangulo)

	//Lista de interfaces
	myInterface := []interface{}{"Hola", 12, 4.90}
	fmt.Println(myInterface...)
	/*
		El mejor momento para usar una interfaz es cuando los structs comparten
		una misma funcion. Se hace mucho mas optimo a nivel de codigo poder interactuar
		con estas
	*/
}
