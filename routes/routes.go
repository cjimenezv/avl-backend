package routes

import (
	"github.com/cjimenezv/avl-backend/controllers"
	"github.com/cjimenezv/avl-backend/middleware"

	"github.com/gin-gonic/gin"
)

// SetupRouter configura todas las rutas de la aplicación
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Grupo de rutas públicas
	public := r.Group("/api")
	{
		public.GET("/public", func(c *gin.Context) {
			c.JSON(200, gin.H{"mensaje": "¡Ruta pública!"})
		})
	}

	// Grupo de rutas protegidas (requiere autenticación JWT con Keycloak)
	protected := r.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	{
		// Ruta de prueba protegida
		protected.GET("/private", func(c *gin.Context) {
			claims, _ := c.Get("claims")
			c.JSON(200, gin.H{
				"mensaje": "Ruta protegida, acceso permitido",
				"claims":  claims,
			})
		})

		// Rutas protegidas del controlador de ubicaciones
		protected.POST("/ubicaciones", controllers.PostUbicacion)
		protected.GET("/ubicaciones/:vehiculoId", controllers.GetUbicaciones)
	}

	return r
}
