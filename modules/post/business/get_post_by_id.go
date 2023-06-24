package business

import (
	"context"
	"dev_community_server/common"
	"dev_community_server/modules/post/entity"
	entity2 "dev_community_server/modules/tag/entity"
)

func (biz *postBusiness) GetPostById(ctx context.Context, id *string) (*entity.Post, error) {
	post, err := biz.postRepo.FindOne(ctx, map[string]interface{}{"id": *id})
	if err != nil {
		return nil, err
	}
	if post == nil {
		return nil, common.NewNotFoundError("Post", common.ErrNotFound)
	}

	author, err := biz.userRepo.FindOne(ctx, map[string]interface{}{"id": post.AuthorId.Hex()})
	post.Author = author

	topic, _ := biz.topicRepo.FindOne(ctx, map[string]interface{}{"id": post.TopicId.Hex()})
	post.Topic = topic

	post.Tags = make([]entity2.Tag, len(post.TagIds))
	for i, id := range post.TagIds {
		tag, _ := biz.tagRepo.FindOne(ctx, map[string]interface{}{"id": id.Hex()})
		post.Tags[i] = *tag
	}

	return post, nil
}
