package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["Jose"] = 14
	m["Pepito"] = 20

	fmt.Println(m)

	//Recorrer map
	for i, v := range m {
		fmt.Println(i, v)
	}

	//Encontrar un valor
	//La variable Ok va a almacenar si existe o no la llave que estamos buscando. Si existe esta variable es TRUE
	value, ok := m["Jose"]
	fmt.Println(value, ok)
}
