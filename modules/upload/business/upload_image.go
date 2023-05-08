package business

import (
	"bytes"
	"context"
)

func (biz *uploadBusiness) UploadImage(ctx context.Context, data []byte) (*string, error) {
	fileBytes := bytes.NewBuffer(data)

	url, err := biz.uploadProvider.Upload(ctx, fileBytes)
	if err != nil {
		return nil, err
	}

	return url, nil
}
