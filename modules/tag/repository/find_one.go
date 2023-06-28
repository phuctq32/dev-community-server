package repository

import (
	"context"
	"dev_community_server/modules/tag/entity"
	"go.mongodb.org/mongo-driver/mongo"
)

func (repo *tagRepository) FindOne(ctx context.Context, filter map[string]interface{}) (*entity.Tag, error) {
	var tag entity.Tag
	if err := repo.tagColl.FindOne(ctx, filter).Decode(&tag); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}

		return nil, err
	}

	return &tag, nil
}
