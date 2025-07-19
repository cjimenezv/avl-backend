package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func CargarVariablesEntorno() {
	appEnv := os.Getenv("APP_ENV")
	if appEnv == "" {
		appEnv = "dev" // Por defecto, usa dev
	}

	archivo := ".env." + appEnv
	err := godotenv.Load(archivo)
	if err != nil {
		log.Fatalf("❌ Error cargando archivo %s: %v", archivo, err)
	}
	fmt.Println("✅ Variables de entorno cargadas desde:", archivo)
}
