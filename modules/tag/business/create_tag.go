package business

import (
	"context"
	"dev_community_server/common"
	"dev_community_server/modules/tag/entity"
	"log"
)

func (biz *tagBusiness) CreateTag(ctx context.Context, data *entity.TagCreate) (*entity.Tag, error) {
	topic, err := biz.topicRepo.FindOne(ctx, map[string]interface{}{"id": data.TopicId})
	if err != nil {
		return nil, err
	}

	if topic == nil {
		return nil, common.NewNotFoundError("Topic", common.ErrNotFound)
	}

	filter, err := common.StructToMap(data)
	if err != nil {
		return nil, err
	}
	log.Println(filter)
	tag, err := biz.tagRepo.FindOne(ctx, filter)
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

	return newTag, nil
}
