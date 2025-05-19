package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Rental struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	UserID     string             `bson:"user_id"`
	CarID      string             `bson:"car_id"`
	Type       string             `bson:"type"` // normal / per_minute
	StartTime  time.Time          `bson:"start_time"`
	EndTime    *time.Time         `bson:"end_time,omitempty"` // nullable
	TotalPrice float64            `bson:"total_price,omitempty"`
}
