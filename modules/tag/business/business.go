package business

import (
	"context"
	"dev_community_server/modules/tag/entity"
	entity2 "dev_community_server/modules/topic/entity"
)

type TagRepository interface {
	Create(ctx context.Context, data *entity.TagCreate) (*entity.Tag, error)
	FindOne(ctx context.Context, filter map[string]interface{}) (*entity.Tag, error)
	Find(ctx context.Context, filter map[string]interface{}) ([]entity.Tag, error)
}

type TopicRepository interface {
	FindOne(ctx context.Context, filter map[string]interface{}) (*entity2.Topic, error)
}

type tagBusiness struct {
	tagRepo   TagRepository
	topicRepo TopicRepository
}

func NewTagBusiness(tagRepo TagRepository, topicRepo TopicRepository) *tagBusiness {
	return &tagBusiness{tagRepo: tagRepo, topicRepo: topicRepo}
}
