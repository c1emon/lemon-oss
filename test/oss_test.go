package test

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/c1emon/lemon_oss/internal/oss/minio"
)

const (
	endpoint        string = ""
	accessKeyID     string = ""
	secretAccessKey string = ""
)

func TestMinio(t *testing.T) {

	client, err := minio.NewMinioClient(endpoint, accessKeyID, secretAccessKey, true)
	if err != nil {
		fmt.Printf("%s", err)
		return
	}

	uri, form, err := client.PostPresignedUrl(context.TODO(), "fxxk", "hello.test")
	if err != nil {
		fmt.Printf("%s", err)
		return
	}

	v, _ := json.Marshal(form)
	fmt.Printf("uri=%s", uri)
	fmt.Printf("form=%s", v)
}
