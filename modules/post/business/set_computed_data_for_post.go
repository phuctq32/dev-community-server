package business

import (
	"context"
	"dev_community_server/common"
	"dev_community_server/modules/post/entity"
	entity2 "dev_community_server/modules/tag/entity"
)

func (biz *postBusiness) SetComputedDataForPostInList(ctx context.Context, post *entity.Post) error {
	// Set author
	if err := biz.SetAuthorData(ctx, post); err != nil {
		return err
	}

	// Set topic
	if err := biz.SetTopicData(ctx, post); err != nil {
		return err
	}

	// Set tags
	if err := biz.SetTagsData(ctx, post); err != nil {
		return err
	}

	// Set count total comments (included replies)
	if err := biz.SetCommentCountData(ctx, post); err != nil {
		return err
	}

	// Calc score
	biz.SetScoreData(ctx, post)

	return nil
}

func (biz *postBusiness) SetComputedData(ctx context.Context, post *entity.Post) error {
	if err := biz.SetComputedDataForPostInList(ctx, post); err != nil {
		return err
	}

	// Set comments (not include replies) and count total comments (included replies)
	if err := biz.SetCommentData(ctx, post); err != nil {
		return err
	}

	return nil
}

func (biz *postBusiness) SetAuthorData(ctx context.Context, post *entity.Post) error {
	userFilter := map[string]interface{}{}
	_ = common.AppendIdQuery(userFilter, "id", post.AuthorId)
	author, err := biz.userRepo.FindOne(ctx, userFilter)
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
	topicFilter := map[string]interface{}{}
	_ = common.AppendIdQuery(topicFilter, "id", post.TopicId)
	topic, err := biz.topicRepo.FindOne(ctx, topicFilter)
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
		tagFilter := map[string]interface{}{}
		_ = common.AppendIdQuery(tagFilter, "id", id)
		tag, err := biz.tagRepo.FindOne(ctx, tagFilter)
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
	cmtFilter := map[string]interface{}{}
	_ = common.AppendIdQuery(cmtFilter, "post_id", *post.Id)
	cmtFilter["parent_comment_id"] = nil
	comments, err := biz.commentRepo.Find(ctx, cmtFilter)
	if err != nil {
		return err
	}

	for i := range comments {
		// Set author data
		userFilter := map[string]interface{}{}
		_ = common.AppendIdQuery(userFilter, "id", comments[i].AuthorId)
		author, err := biz.userRepo.FindOne(ctx, userFilter)
		if err != nil {
			return err
		}
		if author == nil {
			return common.NewNotFoundError("User", common.ErrNotFound)
		}
		comments[i].Author = author

		// Set Reply Count
		repliesFilter := map[string]interface{}{}
		_ = common.AppendIdQuery(repliesFilter, "parent_comment_id", *comments[i].Id)
		replyCount, err := biz.commentRepo.Count(ctx, repliesFilter)
		if err != nil {
			return err
		}
		comments[i].ReplyCount = replyCount

		// Set Score
		score := len(*comments[i].UpVotes) - len(*comments[i].DownVotes)
		comments[i].Score = &score
	}

	post.Comments = comments
	return nil
}

func (biz *postBusiness) SetCommentCountData(ctx context.Context, post *entity.Post) error {
	cmtFilter := map[string]interface{}{}
	_ = common.AppendIdQuery(cmtFilter, "post_id", *post.Id)
	totalCommentCount, err := biz.commentRepo.Count(ctx, cmtFilter)
	if err != nil {
		return err
	}

	post.CommentCount = *totalCommentCount
	return nil
}

func (biz *postBusiness) SetScoreData(ctx context.Context, post *entity.Post) {
	post.Score = len(post.UpVotes) - len(post.DownVotes)
}
