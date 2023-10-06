package ali

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

func NewAliOSS(endpoint, accessId, accessKey string) *AliOSS {

	c, err := oss.New(endpoint, accessId, accessKey)
	if err != nil {
		return nil
	}

	return &AliOSS{
		endpoint:  endpoint,
		accessId:  accessId,
		accessKey: accessKey,
		client:    c,
	}
}

type AliOSS struct {
	endpoint  string
	accessId  string
	accessKey string
	client    *oss.Client
}

func (o *AliOSS) GenSTS(bucketName, objectName string) (string, error) {
	bucket, err := o.client.Bucket(bucketName)
	if err != nil {
		return "", err
	}

	// options can used for set ACL(not tested)
	// should also set callback by options
	signedURL, err := bucket.SignURL(objectName, oss.HTTPPut, 10)
	if err != nil {
		return "", err
	}

	return signedURL, nil
}
