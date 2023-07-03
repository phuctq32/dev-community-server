package business

import (
	"context"
	"dev_community_server/common"
	"dev_community_server/modules/comment/entity"
)

func (biz *commentBusiness) DownVote(ctx context.Context, cmtId string, userId string) (*entity.Comment, error) {
	userFilter := map[string]interface{}{}
	_ = common.AppendIdQuery(userFilter, "id", userId)
	user, _ := biz.userRepo.FindOne(ctx, userFilter)

	cmtFilter := map[string]interface{}{}
	_ = common.AppendIdQuery(cmtFilter, "id", cmtId)
	cmt, err := biz.commentRepo.FindOne(ctx, cmtFilter)
	if err != nil {
		return nil, err
	}
	if cmt == nil {
		return nil, common.NewNotFoundError("Comment", common.ErrNotFound)
	}
	if cmt.ParentCommentId != nil {
		return nil, common.NewCustomBadRequestError("Cannot down vote for a reply")
	}

	updateData := map[string]interface{}{}
	isDownVoting := false
	for i, id := range *cmt.DownVotes {
		if id == *user.Id {
			updateData["down_votes"] = append((*cmt.DownVotes)[:i], (*cmt.DownVotes)[i+1:]...)
			isDownVoting = true
			break
		}
	}
	if !isDownVoting {
		for i, id := range *cmt.UpVotes {
			if id == *user.Id {
				updateData["up_votes"] = append((*cmt.UpVotes)[:i], (*cmt.UpVotes)[i+1:]...)
				break
			}
		}
		updateData["down_votes"] = append(*cmt.DownVotes, *user.Id)
	}

	updatedCmt, err := biz.commentRepo.Update(ctx, cmtFilter, updateData)
	if err != nil {
		return nil, err
	}

	if err = biz.SetComputedDataForCommentInList(ctx, updatedCmt); err != nil {
		return nil, err
	}

	return updatedCmt, nil
}
