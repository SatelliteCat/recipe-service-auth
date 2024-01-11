package app

import (
	"auth/internal/closer"
	"auth/internal/repository"
	userRepository "auth/internal/repository/user"
	"auth/internal/service"
	userService "auth/internal/service/user"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"os"
)

type serviceProvider struct {
	pgPool         *pgxpool.Pool
	userRepository repository.UserRepository

	userService service.UserService
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) PgPool() *pgxpool.Pool {
	if s.pgPool != nil {
		return s.pgPool
	}

	pool, err := pgxpool.New(
		context.Background(),
		fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
			os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"),
			os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD")),
	)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	if err = pool.Ping(context.Background()); err != nil {
		log.Fatalf("failed to ping to database: %v", err)
	}

	closer.Add(func() error {
		pool.Close()
		return nil
	})

	s.pgPool = pool

	return s.pgPool
}

func (s *serviceProvider) UserRepository() repository.UserRepository {
	if s.userRepository == nil {
		s.userRepository = userRepository.NewUserRepository(s.PgPool())
	}

	return s.userRepository
}

func (s *serviceProvider) UserService() service.UserService {
	if s.userService == nil {
		s.userService = userService.NewService(
			s.UserRepository(),
		)
	}

	return s.userService
}
