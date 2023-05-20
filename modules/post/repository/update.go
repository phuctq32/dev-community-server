package repository

import (
	"context"
	"dev_community_server/common"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (repo *postRepository) Update(ctx context.Context, id string, data map[string]interface{}) error {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return common.NewServerError(err)
	}

	if _, err = repo.postColl.UpdateByID(ctx, objId, bson.M{
		"$set": data,
		"$currentDate": bson.M{
			"updated_at": bson.M{"$type": "date"},
		},
	}); err != nil {
		return err
	}

	return nil
}
