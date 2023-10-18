package manager

import (
	"context"
	"crypto/md5"
	"fmt"
	"io"

	"github.com/c1emon/gcommon/gormx"
	"github.com/c1emon/lemon_oss/internal/manager/store"
	"github.com/c1emon/lemon_oss/internal/oss"
	"github.com/c1emon/lemon_oss/internal/oss/minio"
	"github.com/google/uuid"
)

func NewManager() *Manager {
	providerRepo := store.NewGormOSSProviderRepository(gormx.GetGormDB())
	objRepo := store.NewGormOSSObjectRepository(gormx.GetGormDB())

	return &Manager{
		providerRepo: providerRepo,
		objRepo:      objRepo,
		cacher:       store.NewUploadRequestCacher(),
	}
}

type Manager struct {
	providerRepo store.OSSProviderRepository
	objRepo      store.OSSObjectRepository
	cacher       *store.UploadReqCacher
}

func (m *Manager) InitUpload(providerId string, mulitpart bool) string {
	req := &store.UploadRequest{
		Id:         uuid.NewString(),
		ProviderId: providerId,
		Mulitpart:  mulitpart,
		Done:       false,
		ObjectIds:  make([]string, 0),
	}
	m.cacher.Set(req.Id, req)

	return req.Id
}

func (m *Manager) Upload(reqId string, object string, data io.Reader, size int64) error {
	req := m.cacher.Get(reqId)
	if req == nil {
		return fmt.Errorf("invaild upload request id %s", reqId)
	}
	provider, err := m.Find(req.ProviderId)
	if err != nil {
		return fmt.Errorf("bad oss provider id %s", req.ProviderId)
	}

	var client oss.ObjectStorage
	switch provider.Type {
	case store.S3:
		return fmt.Errorf("bad oss provider id %s", req.ProviderId)
	case store.Minio:
		client, err = minio.NewMinioClient(provider.Endpoint, provider.AccessId, provider.AccessKey, provider.Tls)
	case store.Ali:
		return fmt.Errorf("bad oss provider id %s", req.ProviderId)
	case store.COS:
		return fmt.Errorf("bad oss provider id %s", req.ProviderId)
	default:
		return fmt.Errorf("bad oss provider id %s", req.ProviderId)
	}
	if err != nil {
		return fmt.Errorf("failed get oss client %s", provider.Name)
	}

	hasher := md5.New()
	copiedReader := io.TeeReader(data, hasher)
	// hashBuffer := bytes.NewBuffer(nil)

	// pr, pw := io.Pipe()
	// encoder := base64.NewEncoder(base64.StdEncoding, pw)
	// encoderDoneChan := make(chan interface{})

	// go func() {
	// _, err := io.Copy(encoder, copiedReader)
	// if err != nil {
	// 	pw.CloseWithError(err)
	// 	// log.Fatal(err)
	// } else {
	// 	pw.Close()
	// }
	// encoder.Close()
	// hashBuffer.WriteString(fmt.Sprintf("%x", hasher.Sum(nil)))
	// close(encoderDoneChan)
	// }()

	err = client.Upload(context.Background(), provider.BucketName, object, copiedReader, size)
	if err != nil {
		return fmt.Errorf("upload failed for %s", err)
	}
	// hashBuffer.WriteString(fmt.Sprintf("%x", hasher.Sum(nil)))

	record, err := m.RecordUpload(provider.Oid, provider.Id, provider.BucketName, object, fmt.Sprintf("%x", hasher.Sum(nil)))
	if err != nil {
		return fmt.Errorf("upload record failed for %s", err)
	}

	req.ObjectIds = append(req.ObjectIds, record.Id)
	m.cacher.Set(reqId, req)

	return nil
}

func (m *Manager) RecordUpload(oid, providerId, bucket, object, md5 string) (*store.OSSObject, error) {
	obj := &store.OSSObject{
		Bucket:     bucket,
		ObjName:    object,
		ProviderId: providerId,
		MD5:        md5,
		Oid:        oid,
	}
	err := m.objRepo.CreateOne(context.Background(), obj)
	return obj, err
}

func (m *Manager) CompleteUpload(reqId string) (*store.UploadRequest, error) {
	req := m.cacher.Get(reqId)
	if req == nil {
		return nil, fmt.Errorf("invaild upload request id %s", reqId)
	}

	// if mulitPart ...
	req.Done = true

	return req, nil
}

func (m *Manager) Create(p *store.OSSProvider) error {
	return m.providerRepo.CreateOne(context.Background(), p)
}

func (m *Manager) Find(id string) (*store.OSSProvider, error) {
	return m.providerRepo.GetOneById(context.Background(), id)
}
