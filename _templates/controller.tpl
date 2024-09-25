package internal

import (
	"context"
	"upserv/src/apperror"
	"upserv/src/http/request"
	"upserv/src/http/response"
	"upserv/src/service"
	"upserv/src/storage/filter"
	"upserv/src/storage/model"
)

type {{ .ModelName }}Controller struct {
	{{ .LowerModelName }}Service service.I{{ .ModelName }}
}

func (e *{{ .ModelName }}Controller) Get(ctx context.Context, r *request.{{ .ModelName }}Identifier) (*response.{{ .ModelName }}, *apperror.IError) {
	obj, dbErr := e.{{ .LowerModelName }}Service.Get(ctx, r.Id)
	if dbErr != nil {
		return nil, apperror.DBError(ctx, dbErr)
	}
	if obj == nil {
        return nil, apperror.NotFound(ctx)
    }
	res := response.{{ .ModelName }}{}.From(obj)
	return &res, nil
}

func (e *{{ .ModelName }}Controller) Create(ctx context.Context, r *request.{{ .ModelName }}Create) (*response.{{ .ModelName }}, *apperror.IError) {
	{{ .LowerModelName }} := &model.{{ .ModelName }}{
		Title:      r.Title,
	}
	dbErr := e.{{ .LowerModelName }}Service.Create(ctx, {{ .LowerModelName }})
	if dbErr != nil {
		return nil, apperror.DBError(ctx, dbErr)
	}

	resp := response.{{ .ModelName }}{}.From({{ .LowerModelName }})
	return &resp, nil
}

func (e *{{ .ModelName }}Controller) Update(ctx context.Context, r *request.{{ .ModelName }}Update) *apperror.IError {
	{{ .LowerModelName }} := &model.{{ .ModelName }}{}
	{{ .LowerModelName }}.ID = r.Id
	{{ .LowerModelName }}.Title = r.Title
	dbErr := e.{{ .LowerModelName }}Service.Update(ctx, {{ .LowerModelName }})
	if dbErr != nil {
		return apperror.DBError(ctx, dbErr)
	}

	return nil
}

func (e *{{ .ModelName }}Controller) Delete(ctx context.Context, r *request.{{ .ModelName }}Identifier) *apperror.IError {
	{{ .LowerModelName }}, dbErr := e.{{ .LowerModelName }}Service.Get(ctx, r.Id)
	if dbErr != nil {
		return apperror.DBError(ctx, dbErr)
	}
	if {{ .LowerModelName }} == nil {
		return apperror.NotFound(ctx)
	}

	dbErr = e.{{ .LowerModelName }}Service.Delete(ctx, {{ .LowerModelName }}.ID)
	if dbErr != nil {
		return apperror.DBError(ctx, dbErr)
	}
	return nil
}

func (e *{{ .ModelName }}Controller) List(ctx context.Context, r *request.{{ .ModelName }}List) (*response.List, *apperror.IError) {
	f := &filter.Limits{
		Offset: r.Offset,
		Limit:  r.Limit,
	}

	var list []*model.{{ .ModelName }}
	var rowsCount int64
	var dbErr error
	list, rowsCount, dbErr = e.{{ .LowerModelName }}Service.List(ctx, f)
	if dbErr != nil {
		return nil, apperror.DBError(ctx, dbErr)
	}
	result := make([]interface{}, 0)
	for _, val := range list {
		respEvent := response.{{ .ModelName }}{}.From(val)
		result = append(result, respEvent)
	}
	return &response.List{
		Limit:  r.Limit,
		Offset: r.Offset,
		Total:  rowsCount,
		List:   result,
	}, nil
}

func New{{ .ModelName }}Controller({{ .LowerModelName }}Service service.I{{ .ModelName }}) *{{ .ModelName }}Controller {
	return &{{ .ModelName }}Controller{
		{{ .LowerModelName }}Service: {{ .LowerModelName }}Service,
	}
}
