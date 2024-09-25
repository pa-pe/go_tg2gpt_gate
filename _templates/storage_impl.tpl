package internal

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"upserv/src/storage/filter"
	"upserv/src/storage/model"
)

//TODO: move to ../storage.go
//type I{{ .ModelName }} interface {
//	  Get(ctx context.Context, id uint) (*model.{{ .ModelName }}, error)
//	  Create(ctx context.Context, {{ .LowerModelName }} *model.{{ .ModelName }}) error
//	  Update(ctx context.Context, {{ .LowerModelName }} *model.{{ .ModelName }}) error
//	  Delete(ctx context.Context, id uint) error
//	  List(ctx context.Context, filter *filter.Limits) ([]*model.{{ .ModelName }}, int64, error)
//}

type {{ .ModelName }}Impl struct {
	DB *gorm.DB
}

func (f *{{ .ModelName }}Impl) Create(ctx context.Context, {{ .LowerModelName }} *model.{{ .ModelName }}) error {
	db := f.DB.WithContext(ctx)
	err := db.Create({{ .LowerModelName }}).Error
	if err != nil {
		return err
	}
	return nil
}

func (f *{{ .ModelName }}Impl) List(ctx context.Context, filter *filter.Limits) ([]*model.{{ .ModelName }}, int64, error) {
	var {{ .LowerModelName }}s []*model.{{ .ModelName }}

	qb := f.DB.WithContext(ctx).Model(model.{{ .ModelName }}{})
	var rowsCount int64
	err := qb.Count(&rowsCount).Error
	if err != nil {
		return nil, 0, err
	}

	if rowsCount > 0 {
		err = qb.
			Offset(filter.Offset).
			Limit(filter.Limit).Order("created_at ASC, id ASC").Find(&{{ .LowerModelName }}s).Error
		if err != nil {
			return nil, 0, err
		}
	}

	return {{ .LowerModelName }}s, rowsCount, nil
}

func (f *{{ .ModelName }}Impl) Get(ctx context.Context, id uint) (*model.{{ .ModelName }}, error) {
	{{ .LowerModelName }} := &model.{{ .ModelName }}{}
	{{ .LowerModelName }}.ID = id
	db := f.DB.WithContext(ctx)
	dbErr := db.First({{ .LowerModelName }}, {{ .LowerModelName }}).Error

	if dbErr != nil {
		if errors.Is(dbErr, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, dbErr
	}
	return {{ .LowerModelName }}, nil
}

func (f *{{ .ModelName }}Impl) Delete(ctx context.Context, id uint) error {
	e := model.{{ .ModelName }}{}
	e.ID = id
	db := f.DB.WithContext(ctx)
	err := db.Where(e).Delete(&e).Error
	if err != nil {
		return err
	}
	return nil
}

func (f *{{ .ModelName }}Impl) Update(ctx context.Context, {{ .LowerModelName }} *model.{{ .ModelName }}) error {
	db := f.DB.WithContext(ctx)
	err := db.Save({{ .LowerModelName }}).Error
	if err != nil {
		return err
	}
	return nil
}

func New{{ .ModelName }}Storage(db *gorm.DB) *{{ .ModelName }}Impl {
	return &{{ .ModelName }}Impl{
		DB: db,
	}
}
