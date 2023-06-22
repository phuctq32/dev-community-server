package repository

import (
	"context"
	"dev_community_server/modules/tag/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (repo *tagRepository) Create(ctx context.Context, data *entity.TagCreate) (*entity.Tag, error) {
	topicObjId, _ := primitive.ObjectIDFromHex(data.TopicId)
	var tag *entity.Tag = &entity.Tag{
		Name:    data.Name,
		TopicId: topicObjId,
	}

	result, err := repo.tagColl.InsertOne(ctx, tag)
	if err != nil {
		return nil, err
	}
	tag.Id = result.InsertedID.(primitive.ObjectID)

	return tag, nil
}
