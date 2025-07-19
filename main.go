package main

import (
	"github.com/gin-gonic/gin"

	"github.com/cjimenezv/avl-backend/config"
	"github.com/cjimenezv/avl-backend/database"
	"github.com/cjimenezv/avl-backend/routes"
)

func main() {

	// Cargar las variables de entorno
	config.CargarVariablesEntorno()

	// Conectar a MongoDB
	database.ConectarMongoDB()

	router := gin.Default()

	routes.RegisterRoutes(router)

	router.Run(":8080")
}
