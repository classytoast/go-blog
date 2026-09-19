package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"go.uber.org/zap"

	core_logger "github.com/classytoast/go-blog/internal/core/logger"
	core_postgres_pool "github.com/classytoast/go-blog/internal/core/repository/postgres/pool"
	core_http_middleware "github.com/classytoast/go-blog/internal/core/transport/http/middleware"
	core_http_server "github.com/classytoast/go-blog/internal/core/transport/http/server"
	posts_postgres_repository "github.com/classytoast/go-blog/internal/features/posts/repository/postgres"
	posts_service "github.com/classytoast/go-blog/internal/features/posts/service"
	posts_transport_http "github.com/classytoast/go-blog/internal/features/posts/transport/http"
	users_postgres_repository "github.com/classytoast/go-blog/internal/features/users/repository/postgres"
	users_service "github.com/classytoast/go-blog/internal/features/users/service"
	users_transport_http "github.com/classytoast/go-blog/internal/features/users/transport/http"
)

var (
	timeZone = time.UTC
)

func main() {
	time.Local = timeZone

	_ = godotenv.Load()

	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT, syscall.SIGTERM,
	)
	defer cancel()

	logger, err := core_logger.NewLogger(core_logger.NewConfigMust())
	if err != nil {
		fmt.Println("failed to init application logger:", err)
		os.Exit(1)
	}
	defer logger.Close()

	logger.Debug("application time zone", zap.Any("zone", timeZone))

	logger.Debug("initializing postgres connection pool")
	pool, err := core_postgres_pool.NewConnectionPool(
		ctx,
		core_postgres_pool.NewConfigMust(),
	)
	if err != nil {
		logger.Fatal("failed to init postgres connection pool", zap.Error(err))
	}
	defer pool.Close()

	logger.Debug("initializing feature", zap.String("feature", "users"))
	userRepository := users_postgres_repository.NewUserRepository(pool)
	userService := users_service.NewUserService(userRepository)
	userHandler := users_transport_http.NewUserHandler(userService)

	logger.Debug("initializing feature", zap.String("feature", "posts"))
	postRepository := posts_postgres_repository.NewPostRepository(pool)
	postService := posts_service.NewPostService(postRepository)
	postHandler := posts_transport_http.NewPostHTTPHandler(postService)

	logger.Debug("initializing HTTP server")
	httpServer := core_http_server.NewHTTPServer(
		core_http_server.NewConfigMust(),
		logger,
		core_http_middleware.RequestID(),
		core_http_middleware.Logger(logger),
		core_http_middleware.Trace(),
		core_http_middleware.Panic(),
	)
	apiVersionRouter := core_http_server.NewApiVersionRouter(core_http_server.ApiVersion1)
	apiVersionRouter.RegisterRoutes(userHandler.Routes()...)
	apiVersionRouter.RegisterRoutes(postHandler.Routes()...)
	httpServer.RegisterAPIRouters(apiVersionRouter)

	logger.Debug("Starting blog application")

	if err := httpServer.Run(ctx); err != nil {
		logger.Error("HTTP server run error", zap.Error(err))
	}
}
