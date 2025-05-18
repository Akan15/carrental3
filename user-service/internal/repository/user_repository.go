package repository

import (
	"context"

	"github.com/Akan15/carrental3/user-service/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	GetUserByID(id string) (*models.User, error)
}

// Mock implementation
type InMemoryUserRepo struct {
	data map[string]*models.User
}

func NewInMemoryUserRepo() *InMemoryUserRepo {
	return &InMemoryUserRepo{
		data: map[string]*models.User{
			"1": {ID: "1", Name: "Alice", Email: "alice@example.com"},
			"2": {ID: "2", Name: "Bob", Email: "bob@example.com"},
		},
	}
}

func (r *InMemoryUserRepo) GetUserByID(id string) (*models.User, error) {
	user, exists := r.data[id]
	if !exists {
		return nil, nil
	}
	return user, nil
}

type MongoUserRepo struct {
	collection *mongo.Collection
}

func NewMongoUserRepo(db *mongo.Database) *MongoUserRepo {
	return &MongoUserRepo{
		collection: db.Collection("users"),
	}
}

func (r *MongoUserRepo) GetUserByID(id string) (*models.User, error) {
	var user models.User
	err := r.collection.FindOne(context.TODO(), bson.M{"id": id}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
