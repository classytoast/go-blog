package posts_transport_http

import (
	"net/http"

	"github.com/classytoast/go-blog/internal/core/domain"
	core_logger "github.com/classytoast/go-blog/internal/core/logger"
	core_http_request "github.com/classytoast/go-blog/internal/core/transport/http/request"
	core_http_response "github.com/classytoast/go-blog/internal/core/transport/http/response"
)

func (h *PostHTTPHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(logger, w)

	var request CreatePostRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to decode and validate HTTP request",
		)
		return
	}

	postDomain := domain.NewPostUninitialized(
		request.AuthorID,
		request.CategoryID,
		request.Title,
		request.Content,
		request.CoverImageRef,
	)

	postDomain, err := h.postService.CreatePost(ctx, postDomain)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to create post",
		)
		return
	}

	response := CreatePostResponse(postDTOFromDomain(postDomain))

	responseHandler.JSONResponse(response, http.StatusCreated)
}
