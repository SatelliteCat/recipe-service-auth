package service

import (
	"auth/internal/model"
	"context"
)

type UserService interface {
	Create(ctx context.Context, user *model.User) error
	GetByUUID(ctx context.Context, uuid string) (*model.User, error)
}
