package oss

import "github.com/c1emon/lemon_oss/internal/manager/store"

type Upload interface {
	InitUpload(string, string) *store.UploadReq
	GetUploadStatus(string)
	CompleteUpload(string)
}

type SingleUpload interface {
	Upload
	GetPersignedUri() string
}

type MulitpartUpload interface {
	Upload
	GetPersignedUris() string
}
