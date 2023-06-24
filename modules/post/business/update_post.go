package business

import (
	"context"
	"dev_community_server/common"
	"dev_community_server/modules/post/entity"
	entity2 "dev_community_server/modules/tag/entity"
)

func (biz *postBusiness) UpdatePost(ctx context.Context, data *entity.PostUpdate) (*entity.Post, error) {
	post, err := biz.postRepo.FindOne(ctx, map[string]interface{}{"id": *data.Id})
	if err != nil {
		return nil, err
	}
	if post == nil {
		return nil, common.NewNotFoundError("Post", err)
	}

	// Check if user is post's author
	if post.AuthorId.Hex() != data.Author.Id.Hex() {
		return nil, common.NewCustomBadRequestError("User is not author")
	}

	// Get current post's topic
	topic, _ := biz.topicRepo.FindOne(ctx, map[string]interface{}{"id": post.TopicId.Hex()})
	// Find topic if data.TopicId exists
	if data.TopicId != nil {
		existingTopic, err := biz.topicRepo.FindOne(ctx, map[string]interface{}{"id": *data.TopicId})
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
			data.TagIds[i] = tag.Id.Hex()
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
		tag, _ := biz.tagRepo.FindOne(ctx, map[string]interface{}{"id": id.Hex()})
		updatedPost.Tags[i] = *tag
	}

	return updatedPost, nil
}
