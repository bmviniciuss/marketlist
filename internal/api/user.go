package api

import (
	"log/slog"
	"net/http"
)

type UserHandler struct {
	logger *slog.Logger
}

func NewUserHandler(logger *slog.Logger) *UserHandler {
	return &UserHandler{logger: logger}
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("Hanlding user creation")
	w.WriteHeader(http.StatusOK)
}
