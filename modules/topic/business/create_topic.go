package business

import (
	"context"
	"dev_community_server/common"
	"dev_community_server/modules/topic/entity"
)

func (biz *topicBusiness) CreateTopic(ctx context.Context, data *entity.TopicCreate) (*entity.Topic, error) {
	existingTopic, err := biz.repo.FindOne(ctx, map[string]interface{}{"name": data.Name})
	if err != nil {
		return nil, err
	}

	if existingTopic != nil {
		return nil, common.NewCustomBadRequestError("Topic name already exist")
	}

	if data.Description == "" {
		data.Description = "No description"
	}
	topic, err := biz.repo.Create(ctx, data)
	if err != nil {
		return nil, err
	}

	return topic, nil
}
