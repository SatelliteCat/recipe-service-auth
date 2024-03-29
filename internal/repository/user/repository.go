package user

import (
	"auth/internal/client/db"
	"auth/internal/model"
	"auth/internal/repository"
	dbModel "auth/internal/repository/user/model"
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"log/slog"

	sq "github.com/Masterminds/squirrel"
)

const (
	userTableName        = "\"user\""
	userProfileTableName = "user_profile"
	uuidColName          = "uuid"
	createdAtColName     = "created_at"
	updatedAtColName     = "updated_at"
)

var _ repository.UserRepository = (*userRepository)(nil)

type userRepository struct {
	dbPool *pgxpool.Pool
	db     db.Client
}

func NewUserRepository(db db.Client) *userRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) Create(ctx context.Context, user *model.User) error {
	builder := sq.Insert(userTableName).
		Columns(uuidColName, createdAtColName).
		Values(user.UUID, user.CreatedAt).
		PlaceholderFormat(sq.Dollar).
		Suffix("RETURNING " + uuidColName)
	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "query_user_repo_create",
		QueryRaw: query,
	}

	info, err := r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	slog.Debug("inserted", slog.Any("info", info))

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

	q := db.Query{
		Name:     "query_user_repo_get_by_uuid",
		QueryRaw: query,
	}

	var user dbModel.User
	err = r.db.DB().ScanOneContext(ctx, &user, q, args...)
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
