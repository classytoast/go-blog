package users_transport_http

import (
	"net/http"

	"github.com/classytoast/go-blog/internal/core/domain"
	core_logger "github.com/classytoast/go-blog/internal/core/logger"
	core_http_request "github.com/classytoast/go-blog/internal/core/transport/http/request"
	core_http_response "github.com/classytoast/go-blog/internal/core/transport/http/response"
)

func (h *UserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(logger, w)

	logger.Debug("Invoke RegisterUser handler")

	var newUserReq NewUserDTORequest

	if err := core_http_request.DecodeAndValidateRequest(r, &newUserReq); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode and validate HTTP request")
		return
	}

	newUserDomain := domain.NewUserUninitialized{
		Username:  newUserReq.Username,
		Email:     newUserReq.Email,
		Password:  newUserReq.Password,
		AvatarRef: newUserReq.Avatar,
		Bio:       newUserReq.Bio,
	}

	if err := h.userService.RegisterUser(ctx, &newUserDomain); err != nil {
		responseHandler.ErrorResponse(err, "failed to create user")
		return
	}

	responseHandler.JSONResponse(
		map[string]string{"message": "user created"},
		http.StatusCreated,
	)
}
