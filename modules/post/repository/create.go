package repository

import (
	"context"
	"dev_community_server/common"
	"dev_community_server/modules/post/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func (repo *postRepository) Create(ctx context.Context, data *entity.PostCreate) (*entity.Post, error) {
	post := &entity.Post{
		MongoTimestamps: common.MongoTimestamps{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Title:     data.Title,
		Content:   data.Content,
		AuthorId:  *data.Author.Id,
		Images:    data.Images,
		TopicId:   data.TopicId,
		TagIds:    data.TagIds,
		Status:    data.Status,
		UpVotes:   []string{},
		DownVotes: []string{},
		ViewCount: 0,
		IsBlocked: false,
	}
	result, err := repo.postColl.InsertOne(ctx, &post)
	if err != nil {
		return nil, common.NewServerError(err)
	}
	*post.Id = result.InsertedID.(primitive.ObjectID).Hex()

	return post, nil
}
