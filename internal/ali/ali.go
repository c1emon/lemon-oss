package ali

import (
	"fmt"
	"os"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

type OSSType int

type STSProvider interface {
	GenSTS(name string) string
}

type OSS struct {
	Id string

	Provider    OSSType
	Endpoint    string
	AccessKey   string
	AccessId    string
	BucketNames []string
}

type Object struct {
	OSSId string
	Name  string
}

func ali() {
	bucketName := "yourBucketName"
	// yourObjectName填写Object完整路径，完整路径中不能包含Bucket名称
	objectName := "exampledir/exampleobject.txt"

	// 从环境变量中获取访问凭证。运行本代码示例之前，请确保已设置环境变量OSS_ACCESS_KEY_ID和OSS_ACCESS_KEY_SECRET。
	provider, err := oss.NewEnvironmentVariableCredentialsProvider()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	// 创建OSSClient实例。
	// yourEndpoint填写Bucket对应的Endpoint，以华东1（杭州）为例，填写为https://oss-cn-hangzhou.aliyuncs.com。其它Region请按实际情况填写。
	client, err := oss.New("yourEndpoint", "", "", oss.SetCredentialsProvider(&provider))
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	// 获取存储空间。
	bucket, err := client.Bucket(bucketName)
	if err != nil {
	}

	// upload uri
	su, err := bucket.SignURL(objectName, oss.HTTPPut, 10)

	// download uri
	su, err = bucket.SignURL(objectName, oss.HTTPGet, 10)

}
