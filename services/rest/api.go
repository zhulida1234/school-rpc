package rest

import (
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/log"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/zhulida1234/school-rpc/common/httputil"
	"github.com/zhulida1234/school-rpc/config"
	"github.com/zhulida1234/school-rpc/database"
	"github.com/zhulida1234/school-rpc/services/rest/routes"
	"github.com/zhulida1234/school-rpc/services/rest/service"
	"net"
	"strconv"
	"sync/atomic"
	"time"
)

const (
	HealthPath        = "/healthz"
	CreateStudentPath = "/api/v1/create_student"
	UpdateStudentPath = "/api/v1/update_student"
	StudentListPath   = "/api/v1/student_list"
)

type APIConfig struct {
	HTTPServer    config.ServerConfig
	MetricsServer config.ServerConfig
}

type API struct {
	router    *chi.Mux
	apiServer *httputil.HttpServer
	db        *database.SchoolDB
	stopped   atomic.Bool
}

func (a *API) Start(ctx context.Context) error {
	return nil
}

func (a *API) initFromConfig(ctx context.Context, cfg *config.Config) error {
	err := a.initDB(ctx, cfg)
	if err != nil {
		return fmt.Errorf("init db: %w", err)
	}
	a.initRouter(cfg.HttpServer, cfg)
	err = a.startServer(cfg.HttpServer)
	if err != nil {
		return fmt.Errorf("start http server: %w", err)
	}
	return nil
}

func (a *API) initDB(ctx context.Context, cfg *config.Config) error {
	initDB, err := database.NewDB(ctx, cfg.Database)
	if err != nil {
		log.Error("failed to connect to slave database", "err", err)
	}
	a.db = initDB.SchoolDB
	return nil
}

func (a *API) initRouter(server config.ServerConfig, cfg *config.Config) {
	v := new(service.Validator)

	svc := service.NewHandlerSrv(v, a.db)
	apiRouter := chi.NewRouter()
	h := routes.NewRoutes(apiRouter, svc)

	apiRouter.Use(middleware.Timeout(time.Second * 10))
	apiRouter.Use(middleware.Recoverer)
	apiRouter.Use(middleware.Heartbeat(HealthPath))

	apiRouter.Get(fmt.Sprintf(StudentListPath), h.GetStudentList)
	apiRouter.Post(fmt.Sprintf(CreateStudentPath), h.CreateStudent)
	apiRouter.Post(fmt.Sprintf(UpdateStudentPath), h.UpdateStudent)
	a.router = apiRouter
}

func (a *API) Stop(ctx context.Context) error {
	var result error
	if a.apiServer != nil {
		if err := a.apiServer.Stop(ctx); err != nil {
			result = errors.Join(result, fmt.Errorf("stop http server: %w", err))
		}
	}
	if a.db != nil {
		if err := a.db.Close(); err != nil {
			result = errors.Join(result, fmt.Errorf("failed to close DB: %w", err))
		}
	}
	a.stopped.Store(true)
	log.Info("API service shutdown complete")
	return result
}

func (a *API) Stopped() bool {
	return a.stopped.Load()
}

func (a *API) startServer(serverConfig config.ServerConfig) error {
	log.Debug("API server listening...", "port", serverConfig.Port)
	addr := net.JoinHostPort(serverConfig.Host, strconv.Itoa(serverConfig.Port))
	srv, err := httputil.StartHttpServer(addr, a.router)
	if err != nil {
		return fmt.Errorf("failed to start API server: %w", err)
	}
	log.Info("API server started", "addr", srv.Addr().String())
	a.apiServer = srv
	return nil
}

func NewAPI(ctx context.Context, cfg *config.Config) (*API, error) {
	out := &API{}
	err := out.initFromConfig(ctx, cfg)
	if err != nil {
		return nil, errors.Join(err, out.Stop(ctx))
	}
	return out, nil
}
