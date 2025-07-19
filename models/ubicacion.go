package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Ubicacion struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	VehiculoID string             `bson:"vehiculoid" json:"vehiculoId"`
	Lat        float64            `bson:"lat" json:"lat"`
	Lng        float64            `bson:"lng" json:"lng"`
	Timestamp  time.Time          `bson:"timestamp" json:"timestamp"`
}
