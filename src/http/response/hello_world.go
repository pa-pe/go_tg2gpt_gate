package response

import "upserv/src/storage/model"

type HelloWorld struct {
	Cookies
	Title string `json:"title,omitempty" export:"+"`
}

func (t HelloWorld) From(m *model.HelloWorld) HelloWorld {
	t.Title = m.Title
	return t
}
