package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var UbicacionCollection *mongo.Collection

func ConectarMongoDB() {

	// Leer la URI desde la variable de entorno
	uri := os.Getenv("MONGO_URI")
	dbName := os.Getenv("MONGODB_DB")
	collectionName := os.Getenv("MONGODB_COLLECTION")

	if uri == "" {
		log.Fatal("❌ MONGO_URI no está definida en el archivo .env")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("❌ Error conectando a MongoDB: %v", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("❌ Error al hacer ping a MongoDB: %v", err)
	}

	fmt.Println("✅ Conexión a MongoDB exitosa")

	UbicacionCollection = client.Database(dbName).Collection(collectionName)
}
