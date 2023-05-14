package main

import (
	"fmt"
	"io"
	"net/http"
)

type escritorWeb struct{}

func (escritorWeb) Write(p []byte) (int, error) {
	fmt.Println(string(p))
	return 0, nil
}

func main() {

	resp, err := http.Get("http://www.google.com.ar")

	if err != nil {
		fmt.Println(err)
	}

	e := escritorWeb{}
	io.Copy(e, resp.Body)
}
