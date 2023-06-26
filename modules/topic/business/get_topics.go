package business

import (
	"context"
	"dev_community_server/modules/topic/entity"
)

func (biz *topicBusiness) GetTopics(ctx context.Context, filter map[string]interface{}) ([]entity.Topic, error) {
	topics, err := biz.repo.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	return topics, nil
}
