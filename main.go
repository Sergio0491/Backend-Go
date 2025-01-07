package main

import (
	"log"
	"net/http"

	"Backend-go/routes"
)

func main() {
	r := routes.SetupRoutes()

	log.Println("Servidor corriendo en http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
