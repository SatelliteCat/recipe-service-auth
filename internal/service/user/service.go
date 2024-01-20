package user

import (
	"auth/internal/model"
	"auth/internal/repository"
	"auth/internal/service"
	"context"
	"github.com/google/uuid"
	"time"
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
	uuidV7, err := uuid.NewV7()
	if err != nil {
		return err
	}

	user.UUID = uuidV7.String()
	user.CreatedAt = time.Now()

	err = u.userRepository.Create(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (u userService) GetByUUID(ctx context.Context, uuid string) (*model.User, error) {
	user, err := u.userRepository.GetByUUID(ctx, uuid)
	if err != nil {
		return nil, err
	}

	return user, nil
}
