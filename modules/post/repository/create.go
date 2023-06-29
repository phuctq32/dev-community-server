package repository

import (
	"context"
	"dev_community_server/common"
	"dev_community_server/modules/post/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func (repo *postRepository) Create(ctx context.Context, post *entity.Post) (*entity.Post, error) {
	// Convert to object id
	authorOid, err := common.ToObjectId(post.AuthorId)
	if err != nil {
		return nil, err
	}
	topicOid, err := common.ToObjectId(post.TopicId)
	if err != nil {
		return nil, err
	}
	tagOids := make([]primitive.ObjectID, len(post.TagIds))
	for i, tagId := range post.TagIds {
		tagOid, err := common.ToObjectId(tagId)
		if err != nil {
			return nil, err
		}
		tagOids[i] = tagOid
	}

	insertData := &entity.PostInsert{
		MongoTimestamps: common.MongoTimestamps{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Title:     post.Title,
		Content:   post.Content,
		Images:    post.Images,
		AuthorId:  authorOid,
		TopicId:   topicOid,
		TagIds:    tagOids,
		Status:    post.Status,
		UpVotes:   []primitive.ObjectID{},
		DownVotes: []primitive.ObjectID{},
		ViewCount: 0,
		IsBlocked: false,
	}
	result, err := repo.postColl.InsertOne(ctx, &insertData)
	if err != nil {
		return nil, common.NewServerError(err)
	}
	insertedId := result.InsertedID.(primitive.ObjectID).Hex()
	post.Id = &insertedId
	post.CreatedAt = insertData.CreatedAt
	post.UpdatedAt = insertData.UpdatedAt
	post.UpVotes = []string{}
	post.DownVotes = []string{}

	return post, nil
}
