package main

import "fmt"

func main() {
	m1 := make(map[string]int)

	m1["a"] = 8
	m1[9] = "asdasd" //Da error porque el mapa no fue definido asi
	fmt.Println(m1)
}
