package repository

import (
	"context"
	"dev_community_server/common"
	"dev_community_server/modules/post/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (repo *postRepository) Find(ctx context.Context, filter map[string]interface{}) ([]*entity.Post, error) {

	if id, ok := filter["author_id"]; ok {
		objId, err := primitive.ObjectIDFromHex(id.(string))
		if err != nil {
			return nil, err
		}

		filter["author_id"] = objId
	}

	cursor, err := repo.postColl.Find(ctx, filter)
	if err != nil {
		return nil, common.NewServerError(err)
	}

	var posts []*entity.Post
	if err = cursor.All(ctx, &posts); err != nil {
		return nil, common.NewServerError(err)
	}

	return posts, nil
}
