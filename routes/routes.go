package routes

import (
	"github.com/cjimenezv/avl-backend/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.POST("/ubicaciones", controllers.PostUbicacion)
	r.GET("/ubicaciones/:vehiculoId", controllers.GetUbicaciones)
}
