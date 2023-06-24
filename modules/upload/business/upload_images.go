package business

import (
	"bytes"
	"context"
)

func (biz *uploadBusiness) UploadImages(ctx context.Context, data [][]byte) ([]string, error) {
	var res []string

	for _, b := range data {
		fileBytes := bytes.NewBuffer(b)
		url, err := biz.uploadProvider.Upload(ctx, fileBytes)
		if err != nil {
			return nil, err
		}
		res = append(res, *url)
	}

	return res, nil
}
