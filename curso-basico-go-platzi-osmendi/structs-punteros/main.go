package main

import "fmt"

type pc struct {
	ram   int
	disk  int
	brand string
}

func (myPc pc) String() string {
	return fmt.Sprintf("Tengo %d GB RAM, %d GB Disco, y es una %s", myPc.ram, myPc.disk, myPc.brand)
}

func (myPc pc) ping() {
	fmt.Println(myPc.brand, "Pong")
}

func (myPc *pc) duplicateRam() {
	myPc.ram = myPc.ram * 2
}

func main() {
	a := 50
	b := &a //Guarda la direccion de memoria de a

	fmt.Println(a)
	fmt.Println(b)  //Imprime la direccion de memoria de a
	fmt.Println(*b) // Con el * podemos acceder al valor de esa direccion de memoria

	*b = 100
	fmt.Println(a)
	myPc := pc{ram: 16, disk: 200, brand: "msi"}
	fmt.Println(myPc)
	myPc.ping()

	myPc.duplicateRam()
	fmt.Println(myPc.ram)

	fmt.Println(myPc)

}
