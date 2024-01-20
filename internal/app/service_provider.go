package app

import (
	"auth/internal/client/db"
	"auth/internal/client/db/pg"
	"auth/internal/lib/closer"
	"auth/internal/repository"
	userRepository "auth/internal/repository/user"
	"auth/internal/service"
	userService "auth/internal/service/user"
	"auth/internal/transaction"
	"context"
	"fmt"
	"log/slog"
	"os"
)

type serviceProvider struct {
	dbClient       db.Client
	txManager      db.TxManager
	userRepository repository.UserRepository

	userService service.UserService
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) DbClient(ctx context.Context) db.Client {
	if s.dbClient != nil {
		return s.dbClient
	}

	dsn := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"),
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"))
	slog.Debug("dsn", slog.String("dsn", dsn))

	dbClient, err := pg.New(ctx, dsn)
	if err != nil {
		slog.Error("failed to connect to database", slog.String("error", err.Error()))
		panic(err)
	}

	if err = dbClient.DB().Ping(ctx); err != nil {
		slog.Error("failed to ping to database", slog.String("error", err.Error()))
		panic(err)
	}
	slog.Debug("connected to database")

	closer.Add(dbClient.Close)

	s.dbClient = dbClient

	return s.dbClient
}

func (s *serviceProvider) TxManager(ctx context.Context) db.TxManager {
	if s.txManager == nil {
		s.txManager = transaction.NewTransactionManager(s.DbClient(ctx).DB())
	}

	return s.txManager
}

func (s *serviceProvider) UserRepository(ctx context.Context) repository.UserRepository {
	if s.userRepository == nil {
		s.userRepository = userRepository.NewUserRepository(s.DbClient(ctx))
	}

	return s.userRepository
}

func (s *serviceProvider) UserService(ctx context.Context) service.UserService {
	if s.userService == nil {
		s.userService = userService.NewService(
			s.UserRepository(ctx),
		)
	}

	return s.userService
}
