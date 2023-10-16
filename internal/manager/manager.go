package manager

import (
	"context"
	"fmt"

	"github.com/c1emon/gcommon/gormx"
	"github.com/c1emon/lemon_oss/internal/manager/store"
	"github.com/c1emon/lemon_oss/internal/oss/ali"
)

func NewManager() *Manager {
	providerRepo := store.NewGormOSSProviderRepository(gormx.GetGormDB())
	objRepo := store.NewGormOSSObjectRepository(gormx.GetGormDB())

	return &Manager{
		providerRepo: providerRepo,
		objRepo:      objRepo,
	}
}

type Manager struct {
	providerRepo store.OSSProviderRepository
	objRepo      store.OSSObjectRepository
}

func (m *Manager) Create(p *store.OSSProvider) error {
	return m.providerRepo.CreateOne(context.Background(), p)
}

func (m *Manager) Find(id string) (*store.OSSProvider, error) {
	return m.providerRepo.GetOneById(context.Background(), id)
}

func (m *Manager) GenSTS(id, objectName string) (string, error) {
	p, err := m.Find(id)
	if err != nil {
		return "", err
	}

	switch p.Type {
	case store.S3:
		return "", fmt.Errorf("unimplment oss provider type")
	case store.Ali:
		aliOSS := ali.NewAliOSS(p.Endpoint, p.AccessId, p.AccessKey)
		return aliOSS.GenSTS(p.BucketName, objectName)
	default:
		return "", fmt.Errorf("no such oss provider type")
	}
}
