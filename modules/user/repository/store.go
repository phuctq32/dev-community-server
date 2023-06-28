package repository

import (
	userEntity "dev_community_server/modules/user/entity"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	userColl *mongo.Collection
}

func NewUserRepository(db *mongo.Database) *userRepository {
	return &userRepository{userColl: db.Collection(new(userEntity.User).CollectionName())}
}
