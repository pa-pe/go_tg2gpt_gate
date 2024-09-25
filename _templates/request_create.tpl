package request

import (
	"upserv/src/apperror"
)

type {{ .ModelName }}Create struct {
	Title      string `schema:"title,required"`
}

func (s *{{ .ModelName }}Create) InitDefaults() {
}

func (s *{{ .ModelName }}Create) Validate() *apperror.IError {

	return nil
}
