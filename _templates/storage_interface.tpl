
type I{{ .ModelName }} interface {
	Get(ctx context.Context, id uint) (*model.{{ .ModelName }}, error)
	Create(ctx context.Context, {{ .LowerModelName }} *model.{{ .ModelName }}) error
	Update(ctx context.Context, {{ .LowerModelName }} *model.{{ .ModelName }}) error
	Delete(ctx context.Context, id uint) error
	List(ctx context.Context, filter *filter.Limits) ([]*model.{{ .ModelName }}, int64, error)
}