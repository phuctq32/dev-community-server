package repository

import (
	"context"
	"dev_community_server/modules/tag/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (repo *tagRepository) FindOne(ctx context.Context, filter map[string]interface{}) (*entity.Tag, error) {
	if id, ok := filter["id"]; ok {
		objId, err := primitive.ObjectIDFromHex(id.(string))
		if err != nil {
			return nil, err
		}

		filter["_id"] = objId
		delete(filter, "id")
	}

	var tag entity.Tag
	if err := repo.tagColl.FindOne(ctx, filter).Decode(&tag); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}

		return nil, err
	}

	return &tag, nil
}
