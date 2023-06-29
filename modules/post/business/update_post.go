package business

import (
	"context"
	"dev_community_server/common"
	"dev_community_server/modules/post/entity"
	entity2 "dev_community_server/modules/tag/entity"
)

func (biz *postBusiness) UpdatePost(ctx context.Context, data *entity.PostUpdate) (*entity.Post, error) {
	postFilter := map[string]interface{}{}
	if err := common.AppendIdQuery(postFilter, "id", *data.Id); err != nil {
		return nil, err
	}
	post, err := biz.postRepo.FindOne(ctx, postFilter)
	if err != nil {
		return nil, err
	}
	if post == nil {
		return nil, common.NewNotFoundError("Post", err)
	}

	// Check if user is post's author
	if post.AuthorId != *data.Author.Id {
		return nil, common.NewCustomBadRequestError("User is not author")
	}

	// Get current post's topic
	topicFilter := map[string]interface{}{}
	_ = common.AppendIdQuery(topicFilter, "id", post.TopicId)
	topic, _ := biz.topicRepo.FindOne(ctx, topicFilter)
	// Find topic if data.TopicId exists
	if data.TopicId != nil {
		existingTopicFilter := map[string]interface{}{}
		_ = common.AppendIdQuery(existingTopicFilter, "id", *data.TopicId)
		existingTopic, err := biz.topicRepo.FindOne(ctx, existingTopicFilter)
		if err != nil {
			return nil, err
		}
		if existingTopic == nil {
			return nil, common.NewNotFoundError("Topic", common.ErrNotFound)
		}

		// Change post's topic
		topic = existingTopic
	}

	// Find tags if data.TagNames exists
	if len(data.TagNames) > 0 {
		data.TagIds = make([]string, len(data.TagNames))
		for i, tagName := range data.TagNames {
			tagFilter := map[string]interface{}{"name": tagName}
			_ = common.AppendIdQuery(tagFilter, "topic_id", *topic.Id)
			tag, err := biz.tagRepo.FindOne(ctx, tagFilter)
			if err != nil {
				return nil, err
			}
			if tag == nil {
				tag, err = biz.tagRepo.Create(ctx, &entity2.Tag{Name: tagName, TopicId: *topic.Id})
				if err != nil {
					return nil, err
				}
			}
			data.TagIds[i] = *tag.Id
		}
	}

	// Convert data to map
	updateData, _ := common.StructToMap(data)

	updatedPost, err := biz.postRepo.Update(ctx, map[string]interface{}{"id": *data.Id}, updateData)
	if err != nil {
		return nil, err
	}

	updatedPost.Author = data.Author
	updatedPost.Topic = topic

	updatedPost.Tags = make([]entity2.Tag, len(updatedPost.TagIds))
	for i, id := range updatedPost.TagIds {
		tagFilter := map[string]interface{}{}
		_ = common.AppendIdQuery(tagFilter, "id", id)
		tag, _ := biz.tagRepo.FindOne(ctx, tagFilter)
		updatedPost.Tags[i] = *tag
	}

	return updatedPost, nil
}
