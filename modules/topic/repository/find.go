package repository

import (
	"context"
	"dev_community_server/common"
	"dev_community_server/modules/topic/entity"
)

func (repo *topicRepository) Find(ctx context.Context, filter map[string]interface{}) ([]entity.Topic, error) {
	if err := common.ConvertFieldToObjectId(filter, map[string]string{"moderator_ids": "moderator_ids"}); err != nil {
		return nil, common.NewServerError(err)
	}

	cursor, err := repo.topicColl.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	var topics []entity.Topic
	if err := cursor.All(ctx, &topics); err != nil {
		return nil, err
	}

	if topics == nil {
		topics = []entity.Topic{}
	}

	return topics, nil
}
