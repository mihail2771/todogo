package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	core_logger "github.com/mihail2771/todogo/iternal/core/logger"
	core_pgx_pool "github.com/mihail2771/todogo/iternal/core/repository/postgres/pool/pgx"
	core_http_middleware "github.com/mihail2771/todogo/iternal/core/transport/http/middleware"
	core_http_server "github.com/mihail2771/todogo/iternal/core/transport/http/server"
	task_postgres_repository "github.com/mihail2771/todogo/iternal/features/tasks/repositoty/postgres"
	tasks_service "github.com/mihail2771/todogo/iternal/features/tasks/service"
	tasks_transport_http "github.com/mihail2771/todogo/iternal/features/tasks/transport/http"
	users_postgres_repositoty "github.com/mihail2771/todogo/iternal/features/users/repository/postgres"
	users_service "github.com/mihail2771/todogo/iternal/features/users/service"
	users_transport_http "github.com/mihail2771/todogo/iternal/features/users/transport/http"
	"go.uber.org/zap"
)

var (
	timeZone = time.UTC
)

func main() {

	time.Local = timeZone

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

	logger.Debug("application time zone", zap.Any("zone", timeZone))

	logger.Debug("Init connection pool")
	pool, err := core_pgx_pool.NewPool(
		ctx,
		core_pgx_pool.NewMustConfig(),
	)
	if err != nil {
		logger.Fatal("failed init connection pool", zap.Error(err))
	}
	defer pool.Close()

	apiVewrsionRouterV1 := core_http_server.NewAPIVersionRouter(core_http_server.ApiVersion1)

	//USERS
	logger.Debug("Initializing fuature", zap.String("feature", "users"))

	usersRepository := users_postgres_repositoty.NewUserRepository(pool)
	usersService := users_service.NewUserService(usersRepository)
	userTransportHTTP := users_transport_http.NewUserHTTPHandler(usersService)

	apiVewrsionRouterV1.RegistrRoutes(userTransportHTTP.Routers()...)
	//USERS

	//TASKS
	logger.Debug("Initializing fuature", zap.String("feature", "tasks"))

	tasksRepository := task_postgres_repository.NewTasksRepository(pool)
	tasksService := tasks_service.NewTaskService(tasksRepository)
	tasksTransportHTTP := tasks_transport_http.NewTasksHTTPHandler(tasksService)

	apiVewrsionRouterV1.RegistrRoutes(tasksTransportHTTP.Routers()...)
	//TASKS

	// apiVewrsionRouterV2 := core_http_server.NewAPIVersionRouter(core_http_server.ApiVersion2, core_http_middleware.Dummy("v2 ex"))
	// apiVewrsionRouterV2.RegistrRoutes(userRouters...)

	httpServer := core_http_server.NewHTTPServer(
		core_http_server.NewConfigMust(),
		logger,
		core_http_middleware.RequestID(),
		core_http_middleware.Logger(logger),
		core_http_middleware.Trace(),
		core_http_middleware.PanicRecovery(),
	)

	httpServer.RegisterAPIRoutes(
		apiVewrsionRouterV1,
	//	apiVewrsionRouterV2,
	)

	if err := httpServer.Run(ctx); err != nil {
		logger.Error("HTTP server error", zap.Error(err))
	}
}
