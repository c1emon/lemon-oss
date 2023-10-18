package oss

import (
	"context"
	"io"
)

type ObjectStorage interface {
	Upload(ctx context.Context, bucket, object string, data io.Reader, size int64) error
	Download(ctx context.Context, bucket, object string) (io.Reader, error)
}
