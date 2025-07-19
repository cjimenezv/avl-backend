package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/cjimenezv/avl-backend/database"

	"github.com/cjimenezv/avl-backend/models"

	"github.com/gin-gonic/gin"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// POST /ubicaciones
func PostUbicacion(c *gin.Context) {
	var nuevaUbicacion models.Ubicacion

	if err := c.BindJSON(&nuevaUbicacion); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "❌ JSON inválido"})
		return
	}

	// Asignar nuevo ID y timestamp
	nuevaUbicacion.ID = primitive.NewObjectID()
	nuevaUbicacion.Timestamp = time.Now()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := database.UbicacionCollection.InsertOne(ctx, nuevaUbicacion)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "❌ Error insertando ubicación"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"mensaje": "✅ Ubicación guardada exitosamente",
		"id":      result.InsertedID,
	})
}

// GET /ubicaciones/:vehiculoId
func GetUbicaciones(c *gin.Context) {
	vehiculoID := c.Param("vehiculoId")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"vehiculoid": vehiculoID}

	cursor, err := database.UbicacionCollection.Find(ctx, filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "❌ Error consultando ubicaciones"})
		return
	}
	defer cursor.Close(ctx)

	var ubicaciones []models.Ubicacion
	if err := cursor.All(ctx, &ubicaciones); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "❌ Error parseando resultados"})
		return
	}

	c.JSON(http.StatusOK, ubicaciones)
}
