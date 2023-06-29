package business

import (
	"context"
	"dev_community_server/common"
	"dev_community_server/modules/post/entity"
	entity2 "dev_community_server/modules/tag/entity"
)

func (biz *postBusiness) CreatePost(ctx context.Context, data *entity.PostCreate) (*entity.Post, error) {
	if data.Author.GetRoleType() == common.Administrator ||
		data.Author.GetRoleType() == common.Moderator {
		data.Status = entity.Approved
	} else {
		data.Status = entity.Pending
	}
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

	tags := make([]entity2.Tag, len(data.TagNames))

	data.TagIds = make([]string, len(data.TagNames))
	for i, tagName := range data.TagNames {
		tagFilter := map[string]interface{}{"name": tagName}
		_ = common.AppendIdQuery(tagFilter, "topic_id", *topic.Id)
		tag, err := biz.tagRepo.FindOne(ctx, tagFilter)
		if err != nil {
			return nil, err
		}
		if tag == nil {
			tag, err = biz.tagRepo.Create(ctx, &entity2.TagCreate{Name: tagName, TopicId: *topic.Id})
			if err != nil {
				return nil, err
			}
		}
		tagId := *tag.Id
		data.TagIds[i] = tagId

		tags[i] = *tag
	}

	post, err := biz.postRepo.Create(ctx, data)
	if err != nil {
		return nil, err
	}
	post.Author = data.Author
	post.Topic = topic
	post.Tags = tags

	post.Score = 0
	post.CommentCount = 0

	return post, nil
}
