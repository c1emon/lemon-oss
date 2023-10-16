package store

import (
	"context"
	"fmt"

	"github.com/c1emon/gcommon/errorx"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

var _ OSSObjectRepository = &GormOSSObjectRepository{}

func NewGormOSSObjectRepository(db *gorm.DB) *GormOSSObjectRepository {
	r := &GormOSSObjectRepository{
		db: db,
	}
	r.InitDB()
	return r
}

type GormOSSObjectRepository struct {
	db *gorm.DB
}

// CreateOne implements OSSObjectRepository.
func (r *GormOSSObjectRepository) CreateOne(ctx context.Context, obj *OSSObject) error {
	err := r.db.WithContext(ctx).Create(obj).Error
	return errors.Wrap(errorx.From(err), fmt.Sprintf("oss object %s", obj.Name))
}

// DeleteOneById implements OSSObjectRepository.
func (r *GormOSSObjectRepository) DeleteOneById(context.Context, string) error {
	panic("unimplemented")
}

// GetOneById implements OSSObjectRepository.
func (r *GormOSSObjectRepository) GetOneById(ctx context.Context, id string) (*OSSObject, error) {
	obj := &OSSObject{}
	obj.Id = id
	res := r.db.WithContext(ctx).First(obj)
	return obj, errors.Wrap(errorx.From(res.Error), fmt.Sprintf("id %s", id))
}

// InitDB implements OSSObjectRepository.
func (r *GormOSSObjectRepository) InitDB() error {
	return r.db.AutoMigrate(&OSSObject{})
}

// UpdateOneById implements OSSObjectRepository.
func (r *GormOSSObjectRepository) UpdateOneById(context.Context, string, *OSSObject) error {
	panic("unimplemented")
}
