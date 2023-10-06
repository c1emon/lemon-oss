package manager

import (
	"context"
	"fmt"

	"github.com/c1emon/lemon_oss/internal/ali"
	"github.com/c1emon/lemon_oss/pkg/gormx"
)

func NewManager() *Manager {
	providerRepo := &gormOSSProviderRepository{gormx.GetGormDB()}
	providerRepo.InitDB()
	objRepo := &gormOSSObjectRepository{gormx.GetGormDB()}
	objRepo.InitDB()
	return &Manager{
		providerRepo: providerRepo,
		objRepo:      objRepo,
	}
}

type Manager struct {
	providerRepo OSSProviderRepository
	objRepo      OSSObjectRepository
}

func (m *Manager) Create(p *OSSProvider) error {
	return m.providerRepo.CreateOne(context.Background(), p)
}

func (m *Manager) Find(id string) (*OSSProvider, error) {
	return m.providerRepo.GetOneById(context.Background(), id)
}

func (m *Manager) GenSTS(id, objectName string) (string, error) {
	p, err := m.Find(id)
	if err != nil {
		return "", err
	}

	switch p.Type {
	case S3:
		return "", fmt.Errorf("unimplment oss provider type")
	case Ali:
		aliOSS := ali.NewAliOSS(p.Endpoint, p.AccessId, p.AccessKey)
		return aliOSS.GenSTS(p.BucketName, objectName)
	default:
		return "", fmt.Errorf("no such oss provider type")
	}
}
