package main

import (
	"github.com/cjimenezv/avl-backend/config"
	"github.com/cjimenezv/avl-backend/database"
	"github.com/cjimenezv/avl-backend/routes"
)

func main() {
	// Cargar las variables de entorno
	config.CargarVariablesEntorno()

	// Conectar a MongoDB
	database.ConectarMongoDB()

	// Usar el router configurado con rutas
	router := routes.SetupRouter()

	// Correr el servidor
	router.Run(":8080")
}
