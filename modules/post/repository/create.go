package repository

import (
	"context"
	"dev_community_server/common"
	"dev_community_server/modules/post/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (repo *postRepository) Create(ctx context.Context, data *entity.PostCreate) (*entity.Post, error) {
	post := entity.NewPost(data)

	result, err := repo.postColl.InsertOne(ctx, &post)
	if err != nil {
		return nil, common.NewServerError(err)
	}
	post.Id = result.InsertedID.(primitive.ObjectID)

	return post, nil
}
