package repository

import "go.mongodb.org/mongo-driver/mongo"

type tagRepository struct {
	tagColl *mongo.Collection
}

func NewTagRepository(db *mongo.Database) *tagRepository {
	return &tagRepository{tagColl: db.Collection("tags")}
}
