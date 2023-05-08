package upload

import "context"

type UploadProvider interface {
	Upload(ctx context.Context, data interface{}) (*string, error)
}
