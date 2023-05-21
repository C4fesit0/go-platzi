package main

import (
	"fmt"
	"net/http"
)

func HandlerRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}

func HandlerHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Esto es un endpoint")
}
