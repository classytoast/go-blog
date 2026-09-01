package users_transport_http

import (
	"net/http"

	core_logger "github.com/classytoast/go-blog/internal/core/logger"
	core_http_request "github.com/classytoast/go-blog/internal/core/transport/http/request"
	core_http_response "github.com/classytoast/go-blog/internal/core/transport/http/response"
)

type GetUserResponse UserDTOResponse

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(logger, w)

	id, err := core_http_request.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get userID path value",
		)
		return
	}

	userDomain, err := h.userService.GetUser(ctx, id)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get user",
		)
		return
	}

	response := GetUserResponse(userDomain)
	responseHandler.JSONResponse(response, http.StatusOK)
}
