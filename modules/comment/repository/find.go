package repository

import (
	"context"
	"dev_community_server/common"
	"dev_community_server/modules/comment/entity"
)

func (repo *commentRepository) Find(ctx context.Context, filter map[string]interface{}) ([]entity.Comment, error) {
	cursor, err := repo.commentColl.Find(ctx, filter)
	if err != nil {
		return nil, common.NewServerError(err)
	}

	var cmts []entity.Comment
	if err = cursor.All(ctx, &cmts); err != nil {
		return nil, common.NewServerError(err)
	}

	if cmts == nil {
		return []entity.Comment{}, nil
	}

	return cmts, nil
}
