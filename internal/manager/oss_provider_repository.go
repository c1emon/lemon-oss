package manager

import (
	"context"
	"fmt"

	"github.com/c1emon/lemon_oss/pkg/errorx"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

var _ OSSProviderRepository = &gormOSSProviderRepository{}

type gormOSSProviderRepository struct {
	db *gorm.DB
}

// CreateOne implements OSSProviderRepository.
func (r *gormOSSProviderRepository) CreateOne(ctx context.Context, provider *OSSProvider) error {
	err := r.db.WithContext(ctx).Create(provider).Error
	return errors.Wrap(errorx.From(err), fmt.Sprintf("provider %s", provider.Name))
}

// DeleteOneById implements OSSProviderRepository.
func (r *gormOSSProviderRepository) DeleteOneById(context.Context, string) error {
	panic("unimplemented")
}

// GetOneById implements OSSProviderRepository.
func (r *gormOSSProviderRepository) GetOneById(ctx context.Context, id string) (*OSSProvider, error) {
	provider := &OSSProvider{}
	provider.Id = id
	res := r.db.WithContext(ctx).First(provider)
	return provider, errors.Wrap(errorx.From(res.Error), fmt.Sprintf("id %s", id))
}

// InitDB implements OSSProviderRepository.
func (r *gormOSSProviderRepository) InitDB() error {
	return r.db.AutoMigrate(&OSSProvider{})
}

// UpdateOneById implements OSSProviderRepository.
func (r *gormOSSProviderRepository) UpdateOneById(context.Context, string, *OSSProvider) error {
	panic("unimplemented")
}
