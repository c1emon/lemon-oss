package test

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/c1emon/lemon_oss/internal/oss/minio"
)

func TestMinio(t *testing.T) {
	client := minio.NewMinioClient()

	url, form, err := client.PostPresignedUrl(context.Background(), "fxxk", "test.jpg")
	if err != nil {
		fmt.Printf("err=%s", err)
		return
	}
	data, _ := json.Marshal(form)
	fmt.Printf("form=%s\n", data)
	fmt.Printf("url=%s", url)

}
