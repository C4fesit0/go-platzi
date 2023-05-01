package main

import (
	"fmt"
)

func isPalindromo(text string) {
	var textReverse string
	for i := len(text) - 1; i >= 0; i-- {
		textReverse += string(text[i])
	}
	if text == textReverse {
		fmt.Println("Es palindromo")
	} else {
		fmt.Println("No es un palindromo")
	}
}

func main() {

	//Slice
	slice := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(slice, len(slice), cap(slice))
	//Metodos Slice
	fmt.Println(slice[0])
	fmt.Println(slice[:3])
	fmt.Println(slice[2:4])
	fmt.Println(slice[4:])
	//Apend
	slice = append(slice, 7)
	fmt.Println(slice)
	//Apend nueva list
	newSlice := []int{8, 9, 10}
	slice = append(slice, newSlice...)
	//Los tres punto ... desestructuran el newSlice y guarda los valores en slice
	fmt.Println(slice)

	slice2 := []string{"hola", "que", "haces"}
	for i := range slice2 {
		fmt.Println(i)
	}
	isPalindromo("casas")
}
