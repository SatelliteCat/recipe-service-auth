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

type userCreateRequest struct {
}

func CreateUser(ctx context.Context, container di) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req userCreateRequest

		err := render.DecodeJSON(r.Body, &req)
		if errors.Is(err, io.EOF) {
			// Такую ошибку встретим, если получили запрос с пустым телом. Обработаем её отдельно
			slog.Error("r body is empty")
			render.JSON(w, r, struct {
				Status string `json:"status"`
			}{
				Status: "error",
			})
			return
		}
		if err != nil {
			slog.Error("failed to decode request body", sl.Err(err))
			render.JSON(w, r, struct {
				Status string `json:"status"`
			}{
				Status: "error",
			})
			return
		}

		slog.Info("request body decoded", slog.Any("request", req))

		user := &model.User{}
		err = container.UserService(ctx).Create(ctx, user)
		if err != nil {
			slog.Error("failed to create user", sl.Err(err))
			render.JSON(w, r, struct {
				Status string `json:"status"`
			}{
				Status: "error",
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

func GetUser(ctx context.Context, container di) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userUuid := chi.URLParam(r, "uuid")
		user, err := container.UserService(ctx).GetByUUID(ctx, userUuid)
		if err != nil {
			slog.Error("failed to get user", sl.Err(err))
			render.JSON(w, r, struct {
				Status string `json:"status"`
			}{
				Status: "error",
			})
			return
		}

		render.JSON(w, r, struct {
			Status     string `json:"status"`
			Uuid       string `json:"uuid"`
			Created_at string `json:"created_at"`
		}{
			Status:     "OK",
			Uuid:       user.UUID,
			Created_at: user.CreatedAt.String(),
		})
	}
}
