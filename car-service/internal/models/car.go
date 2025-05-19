package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Car struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Brand     string             `bson:"brand" json:"brand"`
	Model     string             `bson:"model" json:"model"`
	City      string             `bson:"city" json:"city"`
	Status    string             `bson:"status" json:"status"` // "free" / "occupied"
	Latitude  float64            `bson:"latitude" json:"latitude"`
	Longitude float64            `bson:"longitude" json:"longitude"`
}
