package app

import (
	"auth/internal/closer"
	"auth/internal/config"
	"context"
	"net/http"
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

	//log := setupLogger(cfg.Env)

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

func (a *App) initConfig(_ context.Context) error {
	config.MustLoad()
	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider()
	return nil
}

//func setupLogger(env string) *slog.Logger {
//	var log *slog.Logger
//
//	switch env {
//	case envLocal:
//		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
//			Level: slog.LevelDebug,
//		}))
//	case envProd:
//		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
//			Level: slog.LevelInfo,
//		}))
//	case envDev:
//		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
//			Level: slog.LevelDebug,
//		}))
//	case envStage:
//		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
//			Level: slog.LevelDebug,
//		}))
//	default:
//		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
//			Level: slog.LevelInfo,
//		}))
//	}
//
//	return log
//}
