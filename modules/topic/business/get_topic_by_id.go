package business

import (
	"context"
	"dev_community_server/common"
	"dev_community_server/modules/topic/entity"
)

func (biz *topicBusiness) GetTopicById(ctx context.Context, id string) (*entity.Topic, error) {
	filter := map[string]interface{}{}
	if err := common.AppendIdQuery(filter, "id", id); err != nil {
		return nil, err
	}
	topic, err := biz.repo.FindOne(ctx, filter)
	if err != nil {
		return nil, err
	}

	if topic == nil {
		return nil, common.NewNotFoundError("Topic", common.ErrNotFound)
	}

	return topic, nil
}
