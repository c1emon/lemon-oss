package manager

import (
	"context"
	"fmt"

	"github.com/c1emon/lemon_oss/pkg/errorx"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

var _ OSSObjectRepository = &gormOSSObjectRepository{}

type gormOSSObjectRepository struct {
	db *gorm.DB
}

// CreateOne implements OSSObjectRepository.
func (r *gormOSSObjectRepository) CreateOne(ctx context.Context, obj *OSSObject) error {
	err := r.db.WithContext(ctx).Create(obj).Error
	return errors.Wrap(errorx.From(err), fmt.Sprintf("oss object %s", obj.Name))
}

// DeleteOneById implements OSSObjectRepository.
func (r *gormOSSObjectRepository) DeleteOneById(context.Context, string) error {
	panic("unimplemented")
}

// GetOneById implements OSSObjectRepository.
func (r *gormOSSObjectRepository) GetOneById(ctx context.Context, id string) (*OSSObject, error) {
	obj := &OSSObject{}
	obj.Id = id
	res := r.db.WithContext(ctx).First(obj)
	return obj, errors.Wrap(errorx.From(res.Error), fmt.Sprintf("id %s", id))
}

// InitDB implements OSSObjectRepository.
func (r *gormOSSObjectRepository) InitDB() error {
	return r.db.AutoMigrate(&OSSObject{})
}

// UpdateOneById implements OSSObjectRepository.
func (r *gormOSSObjectRepository) UpdateOneById(context.Context, string, *OSSObject) error {
	panic("unimplemented")
}
