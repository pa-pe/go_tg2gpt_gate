package response

import "upserv/src/storage/model"

type {{ .ModelName }} struct {
	Id         uint   `json:"id,omitempty"`
	Title      string `json:"title,omitempty"`
}

func (t {{ .ModelName }}) From(o *model.{{ .ModelName }}) {{ .ModelName }} {
	t.Id = o.ID
	t.Title = o.Title

	return t
}
