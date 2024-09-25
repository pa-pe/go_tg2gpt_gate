package request

import "upserv/src/apperror"

type GetHelloWorld struct {
	Cookies
}

func (s *GetHelloWorld) InitDefaults() {}

func (s *GetHelloWorld) Validate() *apperror.IError {
	return nil
}
