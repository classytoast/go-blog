package posts_transport_http

import (
	"net/http"

	"github.com/classytoast/go-blog/internal/core/domain"
	core_logger "github.com/classytoast/go-blog/internal/core/logger"
	core_http_request "github.com/classytoast/go-blog/internal/core/transport/http/request"
	core_http_response "github.com/classytoast/go-blog/internal/core/transport/http/response"
)

func (h *PostHTTPHandler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(logger, w)

	var request CreateCategoryRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to decode and validate HTTP request",
		)
		return
	}

	categoryDomain := domain.NewCategoryUninitialized(request.Name)

	categoryDomain, err := h.postService.CreateCategory(ctx, categoryDomain)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to create category",
		)
		return
	}

	response := CreateCategoryResponse{
		ID:   categoryDomain.ID,
		Name: categoryDomain.Name,
	}

	responseHandler.JSONResponse(response, http.StatusCreated)
}
