package minio

import (
	"context"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

const (
	defaultExpiryTime = time.Second * 24 * 60 * 60 // 1 day

	endpoint        string = "oss.app.clemon.icu:883"
	accessKeyID     string = "R9CwsV40bAL8K2fc"
	secretAccessKey string = "xJyGZToa4qWNZU4NgG79m5DUjELc6URe"
	useSSL          bool   = true
)

type Client struct {
	client *minio.Core
}

func NewMinioClient() *Client {
	client, err := minio.NewCore(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		panic(err)
		// log.Fatalln(err)
	}

	return &Client{
		client: client,
	}
}

func (c *Client) Mulit() {
	c.client.NewMultipartUpload(context.Background(), "", "ss", nil)
}

func (c *Client) PostPresignedUrl(ctx context.Context, bucketName, objectName string) (string, map[string]string, error) {
	expiry := defaultExpiryTime

	policy := minio.NewPostPolicy()
	_ = policy.SetBucket(bucketName)
	_ = policy.SetKey(objectName)
	_ = policy.SetExpires(time.Now().UTC().Add(expiry))

	presignedURL, formData, err := c.client.PresignedPostPolicy(ctx, policy)
	if err != nil {

		// log.Fatalln(err)
		return "", map[string]string{}, err
	}

	return presignedURL.String(), formData, nil
}

func (c *Client) PutPresignedUrl(ctx context.Context, bucketName, objectName string) (string, error) {
	expiry := defaultExpiryTime

	presignedURL, err := c.client.PresignedPutObject(ctx, bucketName, objectName, expiry)
	if err != nil {
		// log.Fatalln(err)
		return "", err
	}

	return presignedURL.String(), nil
}
