package repository

import (
	"context"

	"github.com/Akan15/carrental3/car-service/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type carRepository struct {
	collection *mongo.Collection
}

func NewCarRepository(db *mongo.Database) CarRepository {
	return &carRepository{
		collection: db.Collection("cars"),
	}
}

func (r *carRepository) Create(car *models.Car) (*models.Car, error) {
	res, err := r.collection.InsertOne(context.TODO(), car)
	if err != nil {
		return nil, err
	}
	car.ID = res.InsertedID.(primitive.ObjectID)
	return car, nil
}

func (r *carRepository) GetByID(id string) (*models.Car, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var car models.Car
	err = r.collection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&car)
	return &car, err
}

func (r *carRepository) Update(car *models.Car) error {
	_, err := r.collection.UpdateOne(context.TODO(), bson.M{"_id": car.ID}, bson.M{"$set": car})
	return err
}

func (r *carRepository) Delete(id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = r.collection.DeleteOne(context.TODO(), bson.M{"_id": objID})
	return err
}

func (r *carRepository) List() ([]*models.Car, error) {
	cursor, err := r.collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var cars []*models.Car
	for cursor.Next(context.TODO()) {
		var car models.Car
		if err := cursor.Decode(&car); err == nil {
			cars = append(cars, &car)
		}
	}
	return cars, nil
}

func (r *carRepository) FindByStatus(status string) ([]*models.Car, error) {
	return r.findByFilter(bson.M{"status": status})
}

func (r *carRepository) FindByCity(city string) ([]*models.Car, error) {
	return r.findByFilter(bson.M{"city": city})
}

func (r *carRepository) FindByModel(model string) ([]*models.Car, error) {
	filter := bson.M{"model": bson.M{"$regex": model, "$options": "i"}}
	return r.findByFilter(filter)
}

func (r *carRepository) FindNearby(lat, lon, radius float64) ([]*models.Car, error) {
	filter := bson.M{
		"latitude":  bson.M{"$gt": lat - 0.1, "$lt": lat + 0.1},
		"longitude": bson.M{"$gt": lon - 0.1, "$lt": lon + 0.1},
	}
	return r.findByFilter(filter)
}

func (r *carRepository) ChangeStatus(id, status string) (*models.Car, error) {
	objID, _ := primitive.ObjectIDFromHex(id)
	_, err := r.collection.UpdateOne(context.TODO(), bson.M{"_id": objID}, bson.M{"$set": bson.M{"status": status}})
	if err != nil {
		return nil, err
	}
	return r.GetByID(id)
}

func (r *carRepository) GetLocation(id string) (float64, float64, error) {
	car, err := r.GetByID(id)
	if err != nil {
		return 0, 0, err
	}
	return car.Latitude, car.Longitude, nil
}

func (r *carRepository) UpdateLocation(id string, lat, lon float64) (*models.Car, error) {
	objID, _ := primitive.ObjectIDFromHex(id)
	_, err := r.collection.UpdateOne(context.TODO(), bson.M{"_id": objID}, bson.M{"$set": bson.M{"latitude": lat, "longitude": lon}})
	if err != nil {
		return nil, err
	}
	return r.GetByID(id)
}

func (r *carRepository) findByFilter(filter bson.M) ([]*models.Car, error) {
	cursor, err := r.collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var cars []*models.Car
	for cursor.Next(context.TODO()) {
		var car models.Car
		if err := cursor.Decode(&car); err == nil {
			cars = append(cars, &car)
		}
	}
	return cars, nil
}
