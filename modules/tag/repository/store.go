package repository

import (
	"dev_community_server/modules/tag/entity"
	"go.mongodb.org/mongo-driver/mongo"
)

type tagRepository struct {
	tagColl *mongo.Collection
}

func NewTagRepository(db *mongo.Database) *tagRepository {
	return &tagRepository{tagColl: db.Collection(new(entity.Tag).CollectionName())}
}
