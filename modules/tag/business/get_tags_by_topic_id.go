package business

import (
	"context"
	"dev_community_server/common"
	"dev_community_server/modules/tag/entity"
)

func (biz *tagBusiness) GetTagsByTopicId(ctx context.Context, topicId string) ([]entity.Tag, error) {
	topicFilter := map[string]interface{}{}
	if err := common.AppendIdQuery(topicFilter, "id", topicId); err != nil {
		return nil, err
	}
	topic, err := biz.topicRepo.FindOne(ctx, topicFilter)
	if err != nil {
		return nil, err
	}

	if topic == nil {
		return nil, common.NewNotFoundError("Topic", common.ErrNotFound)
	}

	tagFilter := map[string]interface{}{}
	_ = common.AppendIdQuery(tagFilter, "topic_id", topicId)
	tags, err := biz.tagRepo.Find(ctx, tagFilter)
	if err != nil {
		return nil, err
	}

	for _, tag := range tags {
		tag.Topic = topic
	}

	return tags, nil
}
