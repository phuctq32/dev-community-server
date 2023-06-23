package repository

import (
	"context"
	"dev_community_server/common"
	"dev_community_server/modules/tag/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (repo *tagRepository) Find(ctx context.Context, filter map[string]interface{}) ([]*entity.Tag, error) {
	if id, ok := filter["topic_id"]; ok {
		objId, err := primitive.ObjectIDFromHex(id.(string))
		if err != nil {
			return nil, err
		}
		delete(filter, "topic_id")
		filter["topic_id"] = objId
	}

	cursor, err := repo.tagColl.Find(ctx, filter)
	if err != nil {
		return nil, common.NewServerError(err)
	}

	var tags []*entity.Tag = make([]*entity.Tag, 0)
	if err := cursor.All(ctx, &tags); err != nil {
		return nil, err
	}

	return tags, nil
}
