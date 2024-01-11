package user

import (
	"auth/internal/model"
	"auth/internal/repository"
	"auth/internal/service"
	"context"
)

var _ service.UserService = (*userService)(nil)

type userService struct {
	userRepository repository.UserRepository
}

func NewService(userRepository repository.UserRepository) *userService {
	return &userService{
		userRepository: userRepository,
	}
}

func (u userService) Create(ctx context.Context, user *model.User) error {
	//TODO implement me
	panic("implement me")
}

func (u userService) GetByUUID(ctx context.Context, uuid string) (*model.User, error) {
	//TODO implement me
	panic("implement me")
}
