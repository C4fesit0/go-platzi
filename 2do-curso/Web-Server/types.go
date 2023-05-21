package main

import "net/http"

type Middelware func(http.HandlerFunc) http.HandlerFunc
