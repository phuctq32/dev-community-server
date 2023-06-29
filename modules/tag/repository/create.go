package repository

import (
	"context"
	"dev_community_server/common"
	"dev_community_server/modules/tag/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (repo *tagRepository) Create(ctx context.Context, tag *entity.Tag) (*entity.Tag, error) {
	// Convert to object id
	topicOid, err := common.ToObjectId(tag.TopicId)
	if err != nil {
		return nil, err
	}
	insertData := &entity.TagInsert{
		Name:    tag.Name,
		TopicId: topicOid,
	}

	result, err := repo.tagColl.InsertOne(ctx, &insertData)
	if err != nil {
		return nil, err
	}
	insertedId := result.InsertedID.(primitive.ObjectID).Hex()
	tag.Id = &insertedId

	return tag, nil
}
