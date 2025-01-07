package models

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	openObserveSearchURL string
	openObserveAuthUser  string
	openObserveAuthPass  string
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Advertencia: No se pudo cargar el archivo .env. Usando variables de entorno del sistema.")
	}

	openObserveSearchURL = os.Getenv("OpenObserveSearchURL")
	openObserveAuthUser = os.Getenv("OpenObserveAuthUser")
	openObserveAuthPass = os.Getenv("OpenObserveAuthPass")

	if openObserveSearchURL == "" || openObserveAuthUser == "" || openObserveAuthPass == "" {
		log.Fatal("Faltan variables de entorno necesarias: OpenObserveSearchURL, OpenObserveAuthUser o OpenObserveAuthPass")
	}
}
