package app

import (
	"auth/internal/config"
	"auth/internal/lib/closer"
	"context"
	"io"
	"log/slog"
	"net/http"
	"os"
)

const (
	envLocal = "local"
	envProd  = "prod"
	envDev   = "dev"
	envStage = "stage"
)

type App struct {
	serviceProvider *serviceProvider
	httpServer      *http.Server
}

func NewApp(ctx context.Context) (*App, error) {
	a := &App{}

	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	slog.Info("App initialized")

	return a, nil
}

func (a *App) Run() error {
	defer func() {
		closer.CloseAll()
		closer.Wait()
	}()

	return nil
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
		a.initLogger,
		a.initServiceProvider,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initLogger(_ context.Context) error {
	env := os.Getenv("ENV")

	logFile, err := os.OpenFile("../../var/log.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	closer.Add(func() error {
		return logFile.Close()
	})

	var log *slog.Logger
	loggerOpts := &slog.HandlerOptions{}
	logWriters := []io.Writer{logFile, os.Stdout}

	switch env {
	case envLocal:
		loggerOpts.AddSource = true
		loggerOpts.Level = slog.LevelDebug
	case envProd:
		loggerOpts.AddSource = false
		loggerOpts.Level = slog.LevelInfo
		logWriters = []io.Writer{logFile}
	case envDev:
		loggerOpts.Level = slog.LevelDebug
	case envStage:
		loggerOpts.Level = slog.LevelDebug
	default:
		loggerOpts.AddSource = false
		loggerOpts.Level = slog.LevelInfo
	}

	w := io.MultiWriter(logWriters...)
	log = slog.New(slog.NewJSONHandler(w, loggerOpts))
	slog.SetDefault(log)

	return nil
}

func (a *App) initConfig(_ context.Context) error {
	config.MustLoad()
	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider()
	return nil
}
