package business

import (
	"context"
	"dev_community_server/common"
	"dev_community_server/modules/post/entity"
)

func (biz *postBusiness) UpVote(ctx context.Context, postId string, userId string) (*entity.Post, error) {
	userFilter := map[string]interface{}{}
	_ = common.AppendIdQuery(userFilter, "id", userId)
	currentUser, err := biz.userRepo.FindOne(ctx, userFilter)
	if err != nil {
		return nil, err
	}
	if currentUser == nil {
		return nil, common.NewNotFoundError("User", common.ErrNotFound)
	}

	postFilter := map[string]interface{}{}
	if err = common.AppendIdQuery(postFilter, "id", postId); err != nil {
		return nil, err
	}
	post, err := biz.postRepo.FindOne(ctx, postFilter)
	if err != nil {
		return nil, err
	}
	if post == nil {
		return nil, common.NewNotFoundError("Post", common.ErrNotFound)
	}
	if post.Status == entity.Pending {
		return nil, common.NewCustomBadRequestError("Post is pending")
	}

	updateData := map[string]interface{}{}
	isUpVoting := false
	for i, id := range post.UpVotes {
		if id == *currentUser.Id {
			updateData["up_votes"] = append(post.UpVotes[:i], post.UpVotes[i+1:]...)
			isUpVoting = true
			break
		}
	}
	if !isUpVoting {
		for i, id := range post.DownVotes {
			if id == *currentUser.Id {
				updateData["down_votes"] = append(post.DownVotes[:i], post.DownVotes[i+1:]...)
				break
			}
		}
		updateData["up_votes"] = append(post.UpVotes, *currentUser.Id)
	}

	updatedPost, err := biz.postRepo.Update(ctx, postFilter, updateData)
	if err != nil {
		return nil, err
	}

	if err = biz.SetComputedData(ctx, updatedPost); err != nil {
		return nil, err
	}

	return updatedPost, nil
}
