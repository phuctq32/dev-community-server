package business

import "dev_community_server/components/upload"

type uploadBusiness struct {
	uploadProvider upload.UploadProvider
}

func NewUploadBusiness(uploadProvider upload.UploadProvider) *uploadBusiness {
	return &uploadBusiness{uploadProvider: uploadProvider}
}
