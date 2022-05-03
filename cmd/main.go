package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"identifier/pkg"
	"net/http"
)

func main() {
	router := chi.NewRouter()
	router.Post("/text", pkg.RomanIdentify)

	fmt.Println("server running")
	err := http.ListenAndServe(":8080", router)

	if err != nil {
		panic(err)
	}
}
