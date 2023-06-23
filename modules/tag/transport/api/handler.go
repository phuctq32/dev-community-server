package api

import (
	"context"
	"dev_community_server/components/appctx"
	"dev_community_server/modules/tag/business"
	"dev_community_server/modules/tag/entity"
	"dev_community_server/modules/tag/repository"
	repository2 "dev_community_server/modules/topic/repository"
)

type TagBusiness interface {
	CreateTag(ctx context.Context, data *entity.TagCreate) (*entity.Tag, error)
	GetTagsByTopicId(ctx context.Context, topicId string) ([]*entity.Tag, error)
}

type tagHandler struct {
	appCtx appctx.AppContext
	biz    TagBusiness
}

func NewTagHandler(appCtx appctx.AppContext) *tagHandler {
	tagRepo := repository.NewTagRepository(appCtx.GetAppConfig().GetMongoDbConfig().GetConnection())
	topicRepo := repository2.NewTopicRepository(appCtx.GetAppConfig().GetMongoDbConfig().GetConnection())
	biz := business.NewTagBusiness(tagRepo, topicRepo)

	return &tagHandler{appCtx: appCtx, biz: biz}
}
