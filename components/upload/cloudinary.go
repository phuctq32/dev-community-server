package upload

import (
	"context"
	"dev_community_server/configs"
	"fmt"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"time"
)

type cloudinaryProvider struct {
	cloudinary *cloudinary.Cloudinary
}

func NewCloudinaryProvider(cldConfig configs.CloudinaryConfig) (*cloudinaryProvider, error) {
	cld, err := cloudinary.NewFromParams(
		*cldConfig.GetCloudName(),
		*cldConfig.GetApiKey(),
		*cldConfig.GetApiSecret(),
	)
	if err != nil {
		return nil, err
	}
	return &cloudinaryProvider{cloudinary: cld}, nil
}

func (cld *cloudinaryProvider) Upload(ctx context.Context, file interface{}) (*string, error) {
	uploadRes, err := cld.cloudinary.Upload.Upload(ctx, file, uploader.UploadParams{
		PublicID: fmt.Sprintf("%v", time.Now().Nanosecond()),
		//Folder:   folder,
	})
	if err != nil {
		return nil, err
	}

	return &uploadRes.URL, err
}
