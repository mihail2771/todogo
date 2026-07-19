package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	core_logger "github.com/mihail2771/todogo/iternal/core/logger"
	core_postgres_pool "github.com/mihail2771/todogo/iternal/core/repository/postgres/pool"
	core_http_middleware "github.com/mihail2771/todogo/iternal/core/transport/http/middleware"
	core_http_server "github.com/mihail2771/todogo/iternal/core/transport/http/server"
	users_postgres_repositoty "github.com/mihail2771/todogo/iternal/features/users/repository/postgres"
	users_service "github.com/mihail2771/todogo/iternal/features/users/service"
	users_transport_http "github.com/mihail2771/todogo/iternal/features/users/transport/http"
	"go.uber.org/zap"
)

func main() {

	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)
	defer cancel()

	logger, err := core_logger.NewLogger(core_logger.NewConfigMust())
	if err != nil {
		fmt.Println("faild init logger: ", err)
		os.Exit(1)
	}
	defer logger.Close()

	logger.Debug("Init connection pool")
	pool, err := core_postgres_pool.NewConnectionPool(
		ctx,
		core_postgres_pool.NewMustConfig(),
	)
	if err != nil {
		logger.Fatal("failed init connection pool", zap.Error(err))
	}
	defer pool.Close()

	logger.Debug("Initializing fuature", zap.String("feature", "users"))
	usersRepository := users_postgres_repositoty.NewUserRepository(pool)
	usersService := users_service.NewUserService(usersRepository)

	userTransportHTTP := users_transport_http.NewUserHTTPHandler(usersService)
	userRouters := userTransportHTTP.Routers()

	apiVewrsionRouter := core_http_server.NewAPIVersionRouter(core_http_server.ApiVersion1)
	apiVewrsionRouter.RegistrRoutes(userRouters...)

	httpServer := core_http_server.NewHTTPServer(
		core_http_server.NewConfigMust(),
		logger,
		core_http_middleware.RequestID(),
		core_http_middleware.Logger(logger),
		core_http_middleware.PanicRecovery(),
		core_http_middleware.Trace(),
	)

	httpServer.RegisterAPIRoutes(apiVewrsionRouter)

	if err := httpServer.Run(ctx); err != nil {
		logger.Error("HTTP server error", zap.Error(err))
	}
}
