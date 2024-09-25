package request

import (
	"upserv/src/apperror"
)

type {{ .ModelName }}Update struct {
	Id uint   `schema:"id,required"`
	Title      string `schema:"title,required"`
}

func (s *{{ .ModelName }}Update) InitDefaults() {
}

func (s *{{ .ModelName }}Update) Validate() *apperror.IError {

	return nil
}
