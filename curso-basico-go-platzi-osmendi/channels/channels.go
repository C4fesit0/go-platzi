package main

import "fmt"

func message(text string, c chan<- string) {
	c <- text
}

/*
Con el simbolo <- se le indica que es de entrada de datos. Si ponemos el simbolo
a la izquierda del <-chan significa que es de salida de datos.
*/
func say(text string, c chan<- string) {
	c <- text
}

func main() {
	/*Al crear el channel se le dice el tipo de dato que pasara por el canal
	En este caso es un string
	Despues de la coma, se le indica cuantos datos simultaneos va a manejar es
	canal
	*/
	c := make(chan string, 1)
	fmt.Println("Hello")
	go say("Bye", c)
	fmt.Println(<-c)

	//Range, Close y Select

	d := make(chan string, 2)
	d <- "Mensaje 1"
	d <- "Mensaje 2"

	fmt.Println(len(d), cap(d))

	close(d)

	//d <- "Mensaje 3"

	for message := range d {
		fmt.Println(message)
	}

	//Select
	email1 := make(chan string)
	email2 := make(chan string)
	go message("mensaje1", email1)
	go message("mensaje2", email2)

	for i := 0; i < 2; i++ {
		select {
		case m1 := <-email1:
			fmt.Println("Email recibido de email1", m1)
		case m2 := <-email2:
			fmt.Println("Email recibido de email2", m2)

		}
	}

}
