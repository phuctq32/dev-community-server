package business

import (
	"context"
	"dev_community_server/common"
	"dev_community_server/modules/post/entity"
	entity2 "dev_community_server/modules/tag/entity"
)

func (biz *postBusiness) SetComputedData(ctx context.Context, post *entity.Post) error {
	// Get author
	if err := biz.SetAuthorData(ctx, post); err != nil {
		return err
	}

	// Get topic
	if err := biz.SetTopicData(ctx, post); err != nil {
		return err
	}

	// Get tags
	if err := biz.SetTagsData(ctx, post); err != nil {
		return err
	}

	// Get comments (not include replies) and count total comments (included replies)
	if err := biz.SetCommentData(ctx, post); err != nil {
		return err
	}

	// Calc score
	biz.SetScoreData(ctx, post)

	return nil
}

func (biz *postBusiness) SetAuthorData(ctx context.Context, post *entity.Post) error {
	author, err := biz.userRepo.FindOne(ctx, map[string]interface{}{"id": post.AuthorId.Hex()})
	if err != nil {
		return err
	}
	if author == nil {
		return common.NewNotFoundError("User", common.ErrNotFound)
	}
	post.Author = author
	return nil
}

func (biz *postBusiness) SetTopicData(ctx context.Context, post *entity.Post) error {
	topic, err := biz.topicRepo.FindOne(ctx, map[string]interface{}{"id": post.TopicId.Hex()})
	if err != nil {
		return err
	}
	if topic == nil {
		return common.NewNotFoundError("Topic", common.ErrNotFound)
	}
	post.Topic = topic
	return nil
}

func (biz *postBusiness) SetTagsData(ctx context.Context, post *entity.Post) error {
	post.Tags = make([]entity2.Tag, len(post.TagIds))
	for j, id := range post.TagIds {
		tag, err := biz.tagRepo.FindOne(ctx, map[string]interface{}{"id": id.Hex()})
		if err != nil {
			return err
		}
		if tag == nil {
			return common.NewNotFoundError("Tag", common.ErrNotFound)
		}
		post.Tags[j] = *tag
	}
	return nil
}

func (biz *postBusiness) SetCommentData(ctx context.Context, post *entity.Post) error {
	return nil
}

func (biz *postBusiness) SetScoreData(ctx context.Context, post *entity.Post) {
	post.Score = len(post.UpVotes) - len(post.DownVotes)
}
