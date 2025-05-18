package repository

import (
	"context"
	"time"

	"github.com/Akan15/carrental3/car-service/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoCarRepo struct {
	collection *mongo.Collection
}

func NewMongoCarRepo(db *mongo.Database) *MongoCarRepo {
	return &MongoCarRepo{
		collection: db.Collection("cars"),
	}
}

func (r *MongoCarRepo) GetCarByID(id string) (*models.Car, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var car models.Car
	err := r.collection.FindOne(ctx, bson.M{"id": id}).Decode(&car)
	if err != nil {
		return nil, err
	}
	return &car, nil
}
