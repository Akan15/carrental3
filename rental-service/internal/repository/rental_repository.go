package repository

import (
	"context"
	"time"

	"github.com/Akan15/carrental3/rental-service/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type RentalRepository struct {
	collection *mongo.Collection
}

func NewRentalRepository(db *mongo.Database) *RentalRepository {
	return &RentalRepository{
		collection: db.Collection("rentals"),
	}
}

func (r *RentalRepository) Create(rental *models.Rental) (*models.Rental, error) {
	res, err := r.collection.InsertOne(context.TODO(), rental)
	if err != nil {
		return nil, err
	}
	rental.ID = res.InsertedID.(primitive.ObjectID)
	return rental, nil
}

func (r *RentalRepository) End(id string, endTime time.Time, total float64) (*models.Rental, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": objID}
	update := bson.M{
		"$set": bson.M{
			"end_time":    endTime,
			"total_price": total,
		},
	}
	_, err = r.collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}
	var updated models.Rental
	err = r.collection.FindOne(context.TODO(), filter).Decode(&updated)
	return &updated, err
}

func (r *RentalRepository) GetByID(id string) (*models.Rental, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var rental models.Rental
	err = r.collection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&rental)
	return &rental, err
}

func (r *RentalRepository) List() ([]*models.Rental, error) {
	cursor, err := r.collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var rentals []*models.Rental
	for cursor.Next(context.TODO()) {
		var r models.Rental
		if err := cursor.Decode(&r); err == nil {
			rentals = append(rentals, &r)
		}
	}
	return rentals, nil
}

func (r *RentalRepository) ListByUser(userID string) ([]*models.Rental, error) {
	filter := bson.M{"user_id": userID}
	return r.findByFilter(filter)
}

func (r *RentalRepository) ListByCar(carID string) ([]*models.Rental, error) {
	filter := bson.M{"car_id": carID}
	return r.findByFilter(filter)
}

func (r *RentalRepository) GetActive(userID, carID string) (*models.Rental, error) {
	filter := bson.M{"user_id": userID, "car_id": carID, "end_time": bson.M{"$exists": false}}
	var rental models.Rental
	err := r.collection.FindOne(context.TODO(), filter).Decode(&rental)
	return &rental, err
}

func (r *RentalRepository) Delete(id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = r.collection.DeleteOne(context.TODO(), bson.M{"_id": objID})
	return err
}

func (r *RentalRepository) UpdateType(id, newType string) (*models.Rental, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	_, err = r.collection.UpdateOne(context.TODO(), bson.M{"_id": objID}, bson.M{"$set": bson.M{"type": newType}})
	if err != nil {
		return nil, err
	}
	return r.GetByID(id)
}

func (r *RentalRepository) findByFilter(filter bson.M) ([]*models.Rental, error) {
	cursor, err := r.collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var rentals []*models.Rental
	for cursor.Next(context.TODO()) {
		var rental models.Rental
		if err := cursor.Decode(&rental); err == nil {
			rentals = append(rentals, &rental)
		}
	}
	return rentals, nil
}

func (r *RentalRepository) GetActiveRentals() ([]*models.Rental, error) {
	filter := bson.M{"end_time": bson.M{"$exists": false}}
	return r.findByFilter(filter)
}
