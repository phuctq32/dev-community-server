package api

import (
	"context"
	"dev_community_server/common"
	"dev_community_server/components/appctx"
	"dev_community_server/modules/topic/business"
	"dev_community_server/modules/topic/entity"
	"dev_community_server/modules/topic/repository"
)

type TopicBusiness interface {
	CreateTopic(ctx context.Context, data *entity.TopicCreate) (*entity.Topic, error)
	GetTopics(ctx context.Context, filter common.Filter) ([]*entity.Topic, error)
}

type topicHandler struct {
	appCtx appctx.AppContext
	biz    TopicBusiness
}

func NewTopicHandler(appCtx appctx.AppContext) *topicHandler {
	repo := repository.NewTopicRepository(appCtx.GetAppConfig().GetMongoDbConfig().GetConnection())
	biz := business.NewTopicBusiness(repo)

	return &topicHandler{appCtx: appCtx, biz: biz}
}
