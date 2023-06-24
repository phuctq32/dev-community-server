package business

import (
	"context"
	"dev_community_server/common"
	"dev_community_server/modules/post/entity"
	entity2 "dev_community_server/modules/tag/entity"
)

func (biz *postBusiness) CreatePost(ctx context.Context, data *entity.PostCreate) (*entity.Post, error) {
	topic, err := biz.topicRepo.FindOne(ctx, map[string]interface{}{"id": data.TopicId})
	if err != nil {
		return nil, err
	}
	if topic == nil {
		return nil, common.NewNotFoundError("Topic", common.ErrNotFound)
	}

	tags := make([]entity2.Tag, len(data.TagNames))

	data.TagIds = make([]string, len(data.TagNames))
	for i, tagName := range data.TagNames {
		tag, err := biz.tagRepo.FindOne(ctx, map[string]interface{}{"name": tagName, "topic_id": topic.Id})
		if err != nil {
			return nil, err
		}
		if tag == nil {
			tag, err = biz.tagRepo.Create(ctx, &entity2.TagCreate{Name: tagName, TopicId: topic.Id.Hex()})
			if err != nil {
				return nil, err
			}
		}
		tagId := tag.Id.Hex()
		data.TagIds[i] = tagId

		tags[i] = *tag
	}

	post, err := biz.postRepo.Create(ctx, data)
	if err != nil {
		return nil, err
	}
	post.Topic = topic
	post.Tags = &tags

	return post, nil
}
