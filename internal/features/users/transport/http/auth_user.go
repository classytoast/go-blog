package users_transport_http

// func (h *UserHandler) AuthUser(w http.ResponseWriter, r *http.Request) {
// 	var authUserReq dto.AuthUserRequest

// 	decoder := json.NewDecoder(r.Body)
// 	decoder.DisallowUnknownFields()

// 	if err := decoder.Decode(&authUserReq); err != nil {
// 		http.Error(
// 			w,
// 			"invalid json",
// 			http.StatusBadRequest,
// 		)
// 		return
// 	}

// 	token, err := h.userService.AuthUser(&authUserReq)
// 	if err != nil {
// 		switch {
// 		case errors.Is(err, ers.ErrInvalidAuthData):
// 			http.Error(w, err.Error(), http.StatusBadRequest)

// 		case errors.Is(err, ers.ErrConnectDB):
// 			http.Error(w, err.Error(), http.StatusInternalServerError)

// 		case errors.Is(err, ers.ErrGenerateToken):
// 			http.Error(w, err.Error(), http.StatusInternalServerError)

// 		default:
// 			http.Error(w, "internal server error", http.StatusInternalServerError)
// 		}
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")

// 	json.NewEncoder(w).Encode(map[string]string{
// 		"token": token,
// 	})

// }
