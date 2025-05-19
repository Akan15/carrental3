package repository

import (
	"context"
	"errors"

	"github.com/Akan15/carrental3/user-service/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	GetUserByID(id string) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error) // ✅ добавлено
	CreateUser(user *models.User) error
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
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid user ID format")
	}

	var user models.User
	err = r.collection.FindOne(context.TODO(), bson.M{"_id": objectID}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

func (r *MongoUserRepo) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.collection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *MongoUserRepo) CreateUser(user *models.User) error {
	_, err := r.collection.InsertOne(context.TODO(), user)
	return err
}
