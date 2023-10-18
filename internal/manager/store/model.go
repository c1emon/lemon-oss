package store

import "github.com/c1emon/gcommon/gormx"

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
	Tls        bool    `gorm:"column:tls;type:bool;default:false;not null"`
	Endpoint   string  `gorm:"column:endpoint;type:varchar(256)"`
	AccessId   string  `gorm:"column:access_id;type:varchar(256)"`
	AccessKey  string  `gorm:"column:access_key;type:varchar(256)"`
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
	Name       string `gorm:"column:name;type:varchar(256);not null"`
	ObjName    string `gorm:"column:obj_name;type:varchar(256);uniqueIndex:udx_org_name;not null"`
	Bucket     string `gorm:"column:bucket;type:varchar(256);not null"`
	MD5        string `gorm:"column:md5;type:varchar(256);not null"`
	Oid        string `gorm:"column:oid;type:varchar(256);uniqueIndex:udx_org_name;not null"`
	ProviderId string `gorm:"column:provider_id;type:varchar(256)"`
}

func (OSSObject) TableName() string {
	return "oss_objects"
}

type OSSObjectRepository interface {
	gormx.BaseRepository[OSSObject]
}

type UploadRequest struct {
	Id         string `json:"id"`
	ProviderId string `json:"-"`
	BucketName string `json:"-"`
	ObjName    string `json:"-"`

	ObjectIds []string `json:"object_ids"`
	Mulitpart bool     `json:"-"`
	Done      bool     `json:"done"`
}
