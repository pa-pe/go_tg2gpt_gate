package internal

import (
	"context"
	"strconv"
	"upserv/src/apperror"
	"upserv/src/http/request"
	"upserv/src/http/response"
	"upserv/src/service"
)

type HelloWorldController struct {
	helloWorldService service.IHelloWorld
}

func (e *HelloWorldController) GetHelloWorld(ctx context.Context, r *request.GetHelloWorld) (*response.HelloWorld, *apperror.IError) {
	hw, dbErr := e.helloWorldService.Find(ctx)
	if dbErr != nil {
		return nil, apperror.DBError(ctx, dbErr)
	}
	res := response.HelloWorld{}.From(hw)
	if r.Cookies.Get("hello_world_id") == nil {
		res.Cookies.Set(&response.Cookie{
			Name:  "hello_world_id",
			Value: strconv.FormatUint(uint64(hw.ID), 10),
		})
	}

	return &res, nil
}

func NewHelloWorldController(helloWorldService service.IHelloWorld) *HelloWorldController {
	return &HelloWorldController{
		helloWorldService: helloWorldService,
	}
}
