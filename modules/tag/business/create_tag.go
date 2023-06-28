package business

import (
	"context"
	"dev_community_server/common"
	"dev_community_server/modules/tag/entity"
)

func (biz *tagBusiness) CreateTag(ctx context.Context, data *entity.TagCreate) (*entity.Tag, error) {
	topicFilter := map[string]interface{}{}
	if err := common.AppendIdQuery(topicFilter, "id", data.TopicId); err != nil {
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
	_ = common.AppendIdQuery(tagFilter, "topic_id", *topic.Id)
	tagFilter["name"] = data.Name
	tag, err := biz.tagRepo.FindOne(ctx, tagFilter)
	if err != nil {
		return nil, err
	}

	if tag != nil {
		return nil, common.NewExistingError("Tag")
	}

	newTag, err := biz.tagRepo.Create(ctx, data)
	if err != nil {
		return nil, err
	}
	newTag.Topic = topic

	return newTag, nil
}
