package handler

import (
	"net/http"
	"upserv/src/controller"
	"upserv/src/http/request"
	"upserv/src/http/response"
)

// GetHelloWorld godoc
// @Summary Get hello world test info
// @Router /hello_world [get]
// @Produce json
// @Success 200 {object} response.HelloWorld
// @Failure 400,500 {object} response.errorResp
// @tags test
// @security User
func GetHelloWorld(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	lr := &request.GetHelloWorld{}
	er := request.Extract(lr, r)
	if er != nil {
		response.Error(er, w)
		return
	}
	resp, iErr := controller.ControllerImp.HelloWorldController.GetHelloWorld(ctx, lr)
	if iErr != nil {
		response.Error(iErr, w)
		return
	}

	response.Success(resp, w)
}
