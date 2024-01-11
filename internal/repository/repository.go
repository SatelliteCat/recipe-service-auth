package repository

import (
	"auth/internal/model"
	"context"
)

type UserRepository interface {
	Create(ctx context.Context, user *model.User) error
	GetByUUID(ctx context.Context, uuid string) (*model.User, error)
}
