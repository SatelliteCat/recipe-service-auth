package main

import (
	"auth/internal/app"
	"auth/internal/lib/logger/sl"
	"context"
	"errors"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
)

func main() {
	ctx := context.Background()

	a, err := app.NewApp(ctx)
	if err != nil {
		log.Fatalf("failed to init app: %s", err.Error())
	}

	err = a.Run()
	if err != nil {
		log.Fatalf("failed to run app: %s", err.Error())
	}

	//cfg := config.MustLoad()
	//log := setupLogger(cfg.Env)
	//
	//ctx := context.Background()
	//
	//setupStorage(ctx, cfg, log)
	//
	//router := chi.NewRouter()
	//router.Use(middleware.RequestID)
	//router.Use(mwLogger.New(log))
	//router.Use(middleware.Recoverer)
	//router.Use(middleware.URLFormat)
	//
	//router.Get("/", getUsersHandler)
	//
	//srv := &http.Server{
	//	Addr:         cfg.HttpServer.Address + ":" + cfg.HttpServer.Port,
	//	Handler:      router,
	//	ReadTimeout:  cfg.HttpServer.Timeout,
	//	WriteTimeout: cfg.HttpServer.Timeout,
	//	IdleTimeout:  cfg.HttpServer.IdleTimeout,
	//}
	//
	//listenAndServeWithGracefulShutdown(ctx, srv, log)

	//// Add routes for CRUD operations on users
	//r := http.NewServeMux()
	//r.HandleFunc("/users", createPostHandler)
	//r.HandleFunc("/users/{id}", getPostHandler)
	//r.HandleFunc("/users/{id}", updatePostHandler)
	//r.HandleFunc("/users/{id}", deletePostHandler)
	//
	//// Start the HTTP server
	//srv := &http.Server{
	//	Addr:    ":8080",
	//	Handler: r,
	//}

}

func listenAndServeWithGracefulShutdown(ctx context.Context, srv *http.Server, log *slog.Logger) {
	idleConnsClosed := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)

		log.Info("received an interrupt signal, shutting down", slog.String("signal", (<-sigint).String()))

		// We received an interrupt signal, shut down.
		if err := srv.Shutdown(ctx); err != nil {
			// Error from closing listeners, or context timeout:
			log.Error("HTTP server Shutdown: %v", sl.Err(err))
		}

		close(idleConnsClosed)
	}()

	if err := srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		// Error starting or closing listener:
		log.Error("HTTP server ListenAndServe: %v", sl.Err(err))

		return
	}

	//log.Info("starting HTTP server", "env", cfg.Env, "address", cfg.HttpServer.Address, "port", cfg.HttpServer.Port)
	log.Info("starting HTTP server")

	<-idleConnsClosed
}
