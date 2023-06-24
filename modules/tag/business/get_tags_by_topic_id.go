package business

import (
	"context"
	"dev_community_server/common"
	"dev_community_server/modules/tag/entity"
)

func (biz *tagBusiness) GetTagsByTopicId(ctx context.Context, topicId string) ([]entity.Tag, error) {
	topic, err := biz.topicRepo.FindOne(ctx, map[string]interface{}{"id": topicId})
	if err != nil {
		return nil, err
	}

	if topic == nil {
		return nil, common.NewNotFoundError("Topic", common.ErrNotFound)
	}

	tags, err := biz.tagRepo.Find(ctx, map[string]interface{}{"topic_id": topicId})
	if err != nil {
		return nil, err
	}

	for _, tag := range tags {
		tag.Topic = topic
	}

	return tags, nil
}
