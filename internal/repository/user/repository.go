package user

import (
	"auth/internal/model"
	"auth/internal/repository"
	dbModel "auth/internal/repository/user/model"
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	sq "github.com/Masterminds/squirrel"
)

const (
	userTableName        = "\"user\""
	userProfileTableName = "user_profile"
	uuidColName          = "uuid"
	createdAtColName     = "created_at"
)

var _ repository.UserRepository = (*userRepository)(nil)

type userRepository struct {
	dbPool *pgxpool.Pool
}

func NewUserRepository(dbPool *pgxpool.Pool) *userRepository {
	return &userRepository{
		dbPool: dbPool,
	}
}

func (r *userRepository) Create(_ context.Context, user *model.User) error {

	return nil
}

func (r *userRepository) GetByUUID(ctx context.Context, uuid string) (*model.User, error) {
	builder := sq.Select(uuidColName, createdAtColName).
		From(userTableName).
		Where(sq.Eq{uuidColName: uuid}).
		PlaceholderFormat(sq.Dollar)
	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	var user dbModel.User
	err = r.dbPool.QueryRow(ctx, query, args...).Scan(&user.UUID, &user.CreatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return &model.User{
		UUID:      user.UUID,
		CreatedAt: user.CreatedAt,
	}, nil
}
