package internal

import (
	"context"
	"time"
	"upserv/src/service/cache"
	"upserv/src/storage"
	"upserv/src/storage/filter"
	"upserv/src/storage/model"
)

//TODO: move to ../service.go
//type I{{ .ModelName }} interface {
//	  Get(ctx context.Context, id uint) (*model.{{ .ModelName }}, error)
//	  Create(ctx context.Context, {{ .LowerModelName }} *model.{{ .ModelName }}) error
//	  Update(ctx context.Context, {{ .LowerModelName }} *model.{{ .ModelName }}) error
//	  Delete(ctx context.Context, id uint) error
//    List(ctx context.Context, filter *filter.Limits) ([]*model.{{ .ModelName }}, int64, error)
//}

type {{ .ModelName }}Impl struct {
	cache             cache.ICache
	cacheTime         time.Duration
	cachePrefix       string
	{{ .LowerModelName }}Storage       storage.I{{ .ModelName }}
}

func (u *{{ .ModelName }}Impl) cacheKey(methodName string, params interface{}) string {
	return fmt.Sprintf("__"+u.cachePrefix+"_%s_params_%+v", methodName, params)
}

func (u *{{ .ModelName }}Impl) Get(ctx context.Context, id uint) (*model.{{ .ModelName }}, error) {
	return u.{{ .LowerModelName }}Storage.Get(ctx, id)
}

func (u *{{ .ModelName }}Impl) Create(ctx context.Context, {{ .LowerModelName }} *model.{{ .ModelName }}) error {
	return u.{{ .LowerModelName }}Storage.Create(ctx, {{ .LowerModelName }})
}

func (u *{{ .ModelName }}Impl) Update(ctx context.Context, {{ .LowerModelName }} *model.{{ .ModelName }}) error {
	return u.{{ .LowerModelName }}Storage.Update(ctx, {{ .LowerModelName }})
}

func (u *{{ .ModelName }}Impl) Delete(ctx context.Context, id uint) error {
	return u.{{ .LowerModelName }}Storage.Delete(ctx, id)
}

func (u *{{ .ModelName }}Impl) List(ctx context.Context, filter *filter.Limits) ([]*model.{{ .ModelName }}, int64, error) {
	return u.{{ .LowerModelName }}Storage.List(ctx, filter)
}

func New{{ .ModelName }}Service({{ .LowerModelName }}Storage storage.I{{ .ModelName }}, cache cache.ICache) *{{ .ModelName }}Impl {
	return &{{ .ModelName }}Impl{
		cache:             cache,
		{{ .LowerModelName }}Storage:       {{ .LowerModelName }}Storage,
		cacheTime:         5 * time.Minute,
		cachePrefix:       "{{ .ModelName }}",
	}
}
