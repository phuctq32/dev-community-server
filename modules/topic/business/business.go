package business

import (
	"context"
	"dev_community_server/modules/topic/entity"
)

type TopicRepository interface {
	Create(ctx context.Context, topic *entity.Topic) (*entity.Topic, error)
	FindOne(ctx context.Context, filter map[string]interface{}) (*entity.Topic, error)
	Find(ctx context.Context, filter map[string]interface{}) ([]entity.Topic, error)
}

type topicBusiness struct {
	repo TopicRepository
}

func NewTopicBusiness(repo TopicRepository) *topicBusiness {
	return &topicBusiness{repo: repo}
}
