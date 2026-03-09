package repository

import (
	"context"
	"grpc-user-service-db/config"
	"grpc-user-service-db/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func getCollection() interface{} {
	return config.DB.Collection("users")
}

func CreateUser(user models.User) (models.User, error) {

	collection := config.DB.Collection("users")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, user)

	if err != nil {
		return user, err
	}

	user.ID = result.InsertedID.(primitive.ObjectID)

	return user, nil
}

func GetAllUsers() ([]models.User, error) {

	collection := config.DB.Collection("users")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})

	if err != nil {
		return nil, err
	}

	var users []models.User

	err = cursor.All(ctx, &users)

	return users, err
}
