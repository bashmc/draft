package handlers

import "net/http"

func (h *ApiHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	writeJson(w, 201, Map{"message": "New user has been created"})
}
