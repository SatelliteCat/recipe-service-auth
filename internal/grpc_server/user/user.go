package user

import (
	"auth/internal/model"
	desc "auth/pkg/user_v1"
	"context"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *Implementation) CreateUser(ctx context.Context, req *desc.CreateUserRequest) (*desc.CreateUserResponse, error) {
	user := &model.User{}
	err := s.userService.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	return &desc.CreateUserResponse{
		Uuid: user.UUID,
	}, nil
}

func (s *Implementation) GetUser(ctx context.Context, req *desc.GetUserRequest) (*desc.GetUserResponse, error) {
	user, err := s.userService.GetByUUID(ctx, req.Uuid)
	if err != nil {
		return nil, err
	}

	return &desc.GetUserResponse{
		User: &desc.User{
			Uuid:      user.UUID,
			CreatedAt: timestamppb.New(user.CreatedAt),
		},
	}, nil
}
