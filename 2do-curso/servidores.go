package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	inicio := time.Now()

	servidores := []string{
		"http://platzi.com",
		"http://google.com",
		"http://facebook.com",
		"http://instagram.com",
	}

	for _, servidor := range servidores {
		revisarServidor(servidor)
	}
	tiempo := time.Since(inicio)

	fmt.Println("Tiempo de ejecucion:", tiempo)

}

func revisarServidor(servidor string) {

	_, err := http.Get(servidor)
	if err != nil {
		fmt.Println(servidor, "no esta disponible =(")
	} else {
		fmt.Println(servidor, "esta funcionando normalmente :)")
	}

}
