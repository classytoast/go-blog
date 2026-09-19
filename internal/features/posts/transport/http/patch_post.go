package posts_transport_http

import (
	"net/http"

	core_logger "github.com/classytoast/go-blog/internal/core/logger"
	core_http_request "github.com/classytoast/go-blog/internal/core/transport/http/request"
	core_http_response "github.com/classytoast/go-blog/internal/core/transport/http/response"
)

func (h *PostHTTPHandler) PatchPost(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(logger, w)

	id, err := core_http_request.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(
			err, "failed to get postID path value",
		)
		return
	}

	var request PatchPostRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to decode and validate HTTP request",
		)
		return
	}

	postPatch := postPatchFromRequest(request)

	postDomain, err := h.postService.PatchPost(ctx, id, postPatch)
	if err != nil {
		responseHandler.ErrorResponse(
			err, "failed to patch post",
		)
		return
	}

	response := postDTOFromDomain(postDomain)

	responseHandler.JSONResponse(response, http.StatusOK)
}
