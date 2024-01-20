package user

import (
	"auth/internal/lib/logger/sl"
	"auth/internal/model"
	"auth/internal/service"
	"context"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"io"
	"log/slog"
	"net/http"
)

type di interface {
	UserService(ctx context.Context) service.UserService
}

type handler struct {
	di di
}

func NewHandler(di di) *handler {
	return &handler{
		di: di,
	}
}

type userCreateRequest struct {
}

func (h *handler) CreateUser(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req userCreateRequest

		err := render.DecodeJSON(r.Body, &req)
		if errors.Is(err, io.EOF) {
			// Такую ошибку встретим, если получили запрос с пустым телом. Обработаем её отдельно
			slog.Error("request body is empty")
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, struct {
				Error string `json:"error"`
			}{
				Error: "request body is empty",
			})
			return
		}
		if err != nil {
			slog.Error("failed to decode request body", sl.Err(err))
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, struct {
				Error string `json:"error"`
			}{
				Error: "failed to decode request body",
			})
			return
		}

		slog.Info("request body decoded", slog.Any("request", req))

		user := &model.User{}
		err = h.di.UserService(ctx).Create(ctx, user)
		if err != nil {
			slog.Error("failed to create user", sl.Err(err))
			render.Status(r, http.StatusUnprocessableEntity)
			render.JSON(w, r, struct {
				Error string `json:"error"`
			}{
				Error: "failed to create user",
			})
			return
		}

		render.JSON(w, r, struct {
			Status string `json:"status"`
		}{
			Status: "OK",
		})
	}
}

func (h *handler) GetUser(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userUuid := chi.URLParam(r, "uuid")
		user, err := h.di.UserService(ctx).GetByUUID(ctx, userUuid)
		if err != nil {
			slog.Error("failed to get user", sl.Err(err), slog.String("uuid", userUuid))
			render.Status(r, http.StatusUnprocessableEntity)
			render.JSON(w, r, struct{}{})
			return
		}
		if user == nil {
			render.Status(r, http.StatusNotFound)
			render.JSON(w, r, struct{}{})
			return
		}

		render.JSON(w, r, struct {
			Status    string `json:"status"`
			Uuid      string `json:"uuid"`
			CreatedAt string `json:"created_at"`
		}{
			Status:    "OK",
			Uuid:      user.UUID,
			CreatedAt: user.CreatedAt.String(),
		})
	}
}
