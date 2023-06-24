package api

import (
	"context"
	"dev_community_server/components/appctx"
	"dev_community_server/components/upload"
	"dev_community_server/modules/upload/business"
)

type UploadBusiness interface {
	UploadImage(ctx context.Context, data []byte) (*string, error)
	UploadImages(ctx context.Context, data [][]byte) ([]string, error)
}

type uploadHandler struct {
	business UploadBusiness
}

func NewUploadHandler(appCtx appctx.AppContext) *uploadHandler {
	cldProvider, _ := upload.NewCloudinaryProvider(appCtx.GetCloudinaryConfig())
	biz := business.NewUploadBusiness(cldProvider)

	return &uploadHandler{business: biz}
}
