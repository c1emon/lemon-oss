package store

import "github.com/c1emon/lemon_oss/pkg/gormx"

type OSSType int

const (
	S3 OSSType = iota
	Minio
	Ali
	COS
)

type OSSProvider struct {
	gormx.BaseFields
	Name string `gorm:"column:name;type:varchar(256);uniqueIndex:udx_org_name;not null"`
	Oid  string `gorm:"column:oid;type:varchar(256);uniqueIndex:udx_org_name;not null"`

	Type       OSSType `gorm:"column:oss_type;type:int"`
	Endpoint   string  `gorm:"column:endpoint;type:varchar(256)"`
	AccessKey  string  `gorm:"column:accesskey;type:varchar(256)"`
	AccessId   string  `gorm:"column:accessid;type:varchar(256)"`
	BucketName string  `gorm:"column:bucket_name;type:varchar(256)"`
}

func (OSSProvider) TableName() string {
	return "oss_providers"
}

type OSSProviderRepository interface {
	gormx.BaseRepository[OSSProvider]
}

type OSSObject struct {
	gormx.BaseFields
	Name string `gorm:"column:name;type:varchar(256);uniqueIndex:udx_org_name;not null"`
	Oid  string `gorm:"column:oid;type:varchar(256);uniqueIndex:udx_org_name;not null"`

	ProviderId string `gorm:"column:oss_provider_id;type:varchar(256)"`
}

func (OSSObject) TableName() string {
	return "oss_objects"
}

type OSSObjectRepository interface {
	gormx.BaseRepository[OSSObject]
}

type UploadReq struct {
	Id         string
	ProviderId string
	BucketName string
	ObjName    string

	Mulitpart bool
	Done      bool
}
