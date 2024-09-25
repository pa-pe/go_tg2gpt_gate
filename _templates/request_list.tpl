package request

import (
	"upserv/src/apperror"
)

type {{ .ModelName }}List struct {
	ListParams
}

func (s *{{ .ModelName }}List) InitDefaults() {
	s.Limit = 20
}

func (s *{{ .ModelName }}List) Validate() *apperror.IError {
	return nil
}
