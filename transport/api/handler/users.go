package handler

import (
	"encoding/json"
	"net/http"

	validator "github.com/vincentius93/govalidator"

	"github.com/arymaulanamalik/sicepat_sample/domain/service"
	"github.com/arymaulanamalik/sicepat_sample/pkg/logger"
	"github.com/arymaulanamalik/sicepat_sample/shared"
	"gitlab.sicepat.tech/platform/golib/response"
)

func (h *HandlerImpl) AddUser(r *http.Request) *response.JSONResponse {
	loggerContext := logger.GetLoggerContext(r.Context(), "users", "AddUser")

	var req service.AddUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return response.NewJSONResponse().SetError(response.ErrBadRequest, err.Error())
	}

	if err := validator.Validate(req.Input); err != nil {
		return response.NewJSONResponse().SetError(response.ErrBadRequest, err.Error())
	}

	loggerContext.Printf("%+v", req)

	err := h.Controller.UsersController.AddUser(r.Context(), req)
	if err != nil {
		errRes := shared.GetErrorResponse(err)
		return response.NewJSONResponse().SetError(errRes.Type, errRes.Message).SetLog(errRes.Key, errRes.Message)
	}

	return response.NewJSONResponse().SetMessage(http.StatusText(http.StatusOK))
}
