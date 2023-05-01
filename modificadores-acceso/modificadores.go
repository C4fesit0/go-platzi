package main

import (
	"fmt"
	pk "go-platzi/modificadores-acceso/mypackage"
)

// pk es un alias para el paquete
type car struct {
	brand string
	year  int
}

func main() {
	var myCar pk.CarPublic
	myCar.Brand = "Ferrari"
	myCar.Year = 2020
	fmt.Println(myCar)
	pk.PrintMessage("Hola :D")
	pk.printMessage("Adios")
}
