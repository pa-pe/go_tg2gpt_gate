package handler

import (
	"net/http"
	"upserv/src/controller"
	"upserv/src/http/request"
	"upserv/src/http/response"
)

// Get{{ .ModelName }} godoc
// @Summary Get {{ .ModelName }} info
// @Router /{{ .LowerModelName }}/{id} [get]
// @Param id path integer true "id of {{ .LowerModelName }}"
// @Produce json
// @Success 200 {object} response.{{ .ModelName }}
// @Failure 400,500 {object} response.errorResp
// @tags {{ .ModelName }}
//TODO: swagger description
func Get{{ .ModelName }}(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	lr := &request.{{ .ModelName }}Identifier{}
	er := request.Extract(lr, r)
	if er != nil {
		response.Error(er, w)
		return
	}
	resp, iErr := controller.ControllerImp.{{ .ModelName }}Controller.Get(ctx, lr)
	if iErr != nil {
		response.Error(iErr, w)
		return
	}

	response.Success(resp, w)
}

// Get{{ .ModelName }}List godoc
// @Summary Get {{ .ModelName }}'s list
// @Router /{{ .LowerModelName }}/list [get]
// @Param request.ListParams query request.ListParams false "list params"
// @Produce json
// @Success 200 {object} response.List{List=[]response.{{ .ModelName }}} "List"
// @Failure 400,500 {object} response.errorResp
// @tags {{ .ModelName }}
//TODO: swagger description
func Get{{ .ModelName }}List(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	lr := &request.{{ .ModelName }}List{}
	err := request.Extract(lr, r)
	if err != nil {
		response.Error(err, w)
		return
	}

	resp, iErr := controller.ControllerImp.{{ .ModelName }}Controller.List(ctx, lr)
	if iErr != nil {
		response.Error(iErr, w)
		return
	}
	response.Success(resp, w)
}

// Create{{ .ModelName }} godoc
// @Summary Create {{ .ModelName }}
// @Router /{{ .LowerModelName }} [put]
// @Param request.{{ .ModelName }}Create query request.{{ .ModelName }}Create true "params of {{ .LowerModelName }}"
// @Produce json
// @Success 200 {object} response.{{ .ModelName }}
// @Failure 400,500 {object} response.errorResp
// @tags {{ .ModelName }}
//TODO: swagger description
func Create{{ .ModelName }}(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	lr := &request.{{ .ModelName }}Create{}
	err := request.Extract(lr, r)
	if err != nil {
		response.Error(err, w)
		return
	}
	resp, iErr := controller.ControllerImp.{{ .ModelName }}Controller.Create(ctx, lr)
	if iErr != nil {
		response.Error(iErr, w)
		return
	}
	response.Success(resp, w)
}

// Update{{ .ModelName }} godoc
// @Summary Update {{ .ModelName }}
// @Router /{{ .LowerModelName }} [patch]
// @Param request.{{ .ModelName }}Update query request.{{ .ModelName }}Update true "{{ .LowerModelName }} to update"
// @Produce json
// @Success 202
// @Failure 400,500 {object} response.errorResp
// @tags {{ .ModelName }}
//TODO: swagger description
func Update{{ .ModelName }}(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	lr := &request.{{ .ModelName }}Update{}
	err := request.Extract(lr, r)
	if err != nil {
		response.Error(err, w)
		return
	}
	iErr := controller.ControllerImp.{{ .ModelName }}Controller.Update(ctx, lr)
	if iErr != nil {
		response.Error(iErr, w)
		return
	}
	response.Success(nil, w)
}

// Delete{{ .ModelName }} godoc
// @Summary Delete {{ .ModelName }}
// @Router /{{ .LowerModelName }}/{id} [delete]
// @Param id path integer true "id of {{ .LowerModelName }}"
// @Produce json
// @Success 202
// @Failure 400,500 {object} response.errorResp
// @tags {{ .ModelName }}
//TODO: swagger description
func Delete{{ .ModelName }}(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	lr := &request.{{ .ModelName }}Identifier{}
	err := request.Extract(lr, r)
	if err != nil {
		response.Error(err, w)
		return
	}
	iErr := controller.ControllerImp.{{ .ModelName }}Controller.Delete(ctx, lr)
	if iErr != nil {
		response.Error(iErr, w)
		return
	}
	response.Success(nil, w)
}