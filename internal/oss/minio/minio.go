package minio

import (
	"context"
	"io"
	"time"

	"github.com/c1emon/lemon_oss/internal/oss"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

const (
	defaultExpiryTime = time.Second * 60 * 60 // 1 day
)

var _ oss.ObjectStorage = &Client{}

type Client struct {
	client *minio.Core
}

func NewMinioClient(endpoint, accessKeyID, secretAccessKey string, useSSL bool) (*Client, error) {
	client, err := minio.NewCore(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		return nil, err
	}

	return &Client{
		client: client,
	}, nil
}

func (c *Client) Upload(ctx context.Context, bucket, object string, data io.Reader, size int64) error {
	_, err := c.client.Client.PutObject(ctx, bucket, object, data, size, minio.PutObjectOptions{ContentType: "application/octet-stream"})
	return err
}

func (c *Client) Download(ctx context.Context, bucket, object string) (io.Reader, error) {
	return c.client.Client.GetObject(ctx, bucket, object, minio.GetObjectOptions{})
}

func (c *Client) PostPresignedUrl(ctx context.Context, bucketName, objectName string) (string, map[string]string, error) {

	policy := minio.NewPostPolicy()
	_ = policy.SetSuccessActionRedirect("http://baidu.com")
	_ = policy.SetBucket(bucketName)
	_ = policy.SetKey(objectName)
	_ = policy.SetExpires(time.Now().UTC().Add(defaultExpiryTime))

	presignedURL, formData, err := c.client.PresignedPostPolicy(ctx, policy)
	if err != nil {
		return "", nil, err
	}

	return presignedURL.String(), formData, nil
}

func (c *Client) PutPresignedUrl(ctx context.Context, bucketName, objectName string) (string, error) {

	presignedURL, err := c.client.PresignedPutObject(ctx, bucketName, objectName, defaultExpiryTime)
	if err != nil {
		return "", err
	}

	return presignedURL.String(), nil
}
