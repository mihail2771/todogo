package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	core_config "github.com/mihail2771/todogo/iternal/core/config"
	core_logger "github.com/mihail2771/todogo/iternal/core/logger"
	core_pgx_pool "github.com/mihail2771/todogo/iternal/core/repository/postgres/pool/pgx"
	core_http_middleware "github.com/mihail2771/todogo/iternal/core/transport/http/middleware"
	core_http_server "github.com/mihail2771/todogo/iternal/core/transport/http/server"
	statistics_postgres_repositoty "github.com/mihail2771/todogo/iternal/features/statistics/repository/postgres"
	statistics_service "github.com/mihail2771/todogo/iternal/features/statistics/service"
	statistics_transport_http "github.com/mihail2771/todogo/iternal/features/statistics/transport/http"
	task_postgres_repository "github.com/mihail2771/todogo/iternal/features/tasks/repositoty/postgres"
	tasks_service "github.com/mihail2771/todogo/iternal/features/tasks/service"
	tasks_transport_http "github.com/mihail2771/todogo/iternal/features/tasks/transport/http"
	users_postgres_repositoty "github.com/mihail2771/todogo/iternal/features/users/repository/postgres"
	users_service "github.com/mihail2771/todogo/iternal/features/users/service"
	users_transport_http "github.com/mihail2771/todogo/iternal/features/users/transport/http"
	"go.uber.org/zap"

	_ "github.com/mihail2771/todogo/docs"
)

// @title 			Golang Todo API
// @version 		1.0
// @description 	Todo Application RES-API scheme
// @host 			localhost:5050
// @schemes 		http
// @BasePath 		/api/v1
func main() {
	cfg := core_config.NewConfigMust()
	time.Local = cfg.TimeZone

	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)
	defer cancel()

	logger, err := core_logger.NewLogger(core_logger.NewConfigMust())
	if err != nil {
		fmt.Println("failed init logger: ", err)
		os.Exit(1)
	}
	defer logger.Close()

	logger.Debug("application time zone", zap.Any("zone", time.Local))

	logger.Debug("Init connection pool")
	pool, err := core_pgx_pool.NewPool(
		ctx,
		core_pgx_pool.NewMustConfig(),
	)
	if err != nil {
		logger.Fatal("failed init connection pool", zap.Error(err))
	}
	defer pool.Close()

	apiVersionRouterV1 := core_http_server.NewAPIVersionRouter(core_http_server.ApiVersion1)

	//USERS
	logger.Debug("Initializing feature", zap.String("feature", "users"))

	usersRepository := users_postgres_repositoty.NewUserRepository(pool)
	usersService := users_service.NewUserService(usersRepository)
	userTransportHTTP := users_transport_http.NewUserHTTPHandler(usersService)

	apiVersionRouterV1.RegisterRoutes(userTransportHTTP.Routers()...)
	//USERS

	//TASKS
	logger.Debug("Initializing feature", zap.String("feature", "tasks"))

	tasksRepository := task_postgres_repository.NewTasksRepository(pool)
	tasksService := tasks_service.NewTaskService(tasksRepository)
	tasksTransportHTTP := tasks_transport_http.NewTasksHTTPHandler(tasksService)

	apiVersionRouterV1.RegisterRoutes(tasksTransportHTTP.Routers()...)
	//TASKS

	//STATISTICS
	logger.Debug("Initializing feature", zap.String("feature", "statistics"))

	statisticsRepository := statistics_postgres_repositoty.NewStatisticsRepository(pool)
	statisticsService := statistics_service.NewStatisticsService(statisticsRepository)
	statisticsTransportHTTP := statistics_transport_http.NewStatisticHTTPHandler(statisticsService)

	apiVersionRouterV1.RegisterRoutes(statisticsTransportHTTP.Routers()...)
	//STATISTICS

	// apiVersionRouterV2 := core_http_server.NewAPIVersionRouter(core_http_server.ApiVersion2, core_http_middleware.Dummy("v2 ex"))
	// apiVersionRouterV2.RegisterRoutes(userRouters...)

	httpServer := core_http_server.NewHTTPServer(
		core_http_server.NewConfigMust(),
		logger,
		core_http_middleware.CORS(),
		core_http_middleware.RequestID(),
		core_http_middleware.Logger(logger),
		core_http_middleware.Trace(),
		core_http_middleware.PanicRecovery(),
	)

	httpServer.RegisterAPIRoutes(
		apiVersionRouterV1,
	//	apiVersionRouterV2,
	)

	httpServer.RegisterSwagger()

	if err := httpServer.Run(ctx); err != nil {
		logger.Error("HTTP server error", zap.Error(err))
	}
}
