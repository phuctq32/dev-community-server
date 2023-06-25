package business

import (
	"context"
	"dev_community_server/common"
	"dev_community_server/modules/post/entity"
	"errors"
)

func (biz *postBusiness) BlockPost(ctx context.Context, postId *string, user *common.Requester) (*entity.Post, error) {
	post, err := biz.postRepo.FindOne(ctx, map[string]interface{}{"id": *postId})
	if err != nil {
		return nil, err
	}
	if post == nil {
		return nil, common.NewNotFoundError("Post", common.ErrNotFound)
	}

	topic, _ := biz.topicRepo.FindOne(ctx, map[string]interface{}{"id": post.TopicId.Hex()})

	// Check role: user can be mod or admin
	// If user is admin -> auto allow, if user is a mod -> check topic's mods contains user
	if (*user).GetRoleType() == common.Moderator {
		isValid := false
		for _, modId := range topic.ModeratorIds {
			if modId.Hex() == (*user).GetUserId() {
				isValid = true
				break
			}
		}

		if !isValid {
			return nil, common.NewNoPermissionError(errors.New("User is not topic's moderator"))
		}
	}

	if post.IsBlocked {
		return nil, common.NewCustomBadRequestError("Post already blocked before")
	}

	updatedPost, err := biz.postRepo.Update(
		ctx,
		map[string]interface{}{"id": *postId},
		map[string]interface{}{"is_blocked": true},
	)

	if err = biz.SetComputedData(ctx, updatedPost); err != nil {
		return nil, err
	}

	return updatedPost, nil
}
