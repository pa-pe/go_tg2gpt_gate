package request

import (
	"upserv/src/apperror"
)

type {{ .ModelName }}Identifier struct {
	Id uint `schema:"id,required"`
}

func (s *{{ .ModelName }}Identifier) InitDefaults() {
}

func (s *{{ .ModelName }}Identifier) Validate() *apperror.IError {
	return nil
}
