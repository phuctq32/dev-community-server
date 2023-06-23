package business

import (
	"context"
	"dev_community_server/common"
	"dev_community_server/modules/topic/entity"
)

func (biz *topicBusiness) GetTopicById(ctx context.Context, id string) (*entity.Topic, error) {
	topic, err := biz.repo.FindOne(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return nil, err
	}

	if topic == nil {
		return nil, common.NewNotFoundError("Topic", common.ErrNotFound)
	}

	return topic, nil
}
