package repository

import "go.mongodb.org/mongo-driver/mongo"

type userRepository struct {
	userColl *mongo.Collection
}

func NewUserRepository(db *mongo.Database) *userRepository {
	return &userRepository{userColl: db.Collection("users")}
}
