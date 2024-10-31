package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
	"github.com/urfave/cli/v2"
	"github.com/zhulida1234/school-rpc/common/cliapp"
	"github.com/zhulida1234/school-rpc/common/opio"
	"github.com/zhulida1234/school-rpc/config"
	"github.com/zhulida1234/school-rpc/database"
	flags2 "github.com/zhulida1234/school-rpc/flags"
	"github.com/zhulida1234/school-rpc/services/rest"
	"github.com/zhulida1234/school-rpc/services/rpc"
	"strconv"
)

func runRpc(ctx *cli.Context, causeFunc context.CancelCauseFunc) (cliapp.Lifecycle, error) {
	fmt.Println("running grpc server...")
	cfg := config.NewConfig(ctx)
	grpcServerCfg := &rpc.RpcServerConfig{
		GrpcHost: cfg.RpcServer.Host,
		GrpcPort: strconv.Itoa(cfg.RpcServer.Port),
	}
	db, err := database.NewDB(ctx.Context, cfg.Database)
	if err != nil {
		log.Error("failed to connect to database", "err", err)
		return nil, err
	}
	return rpc.NewRpcServer(grpcServerCfg, db)
}

func runRestApi(ctx *cli.Context, shutdown context.CancelCauseFunc) (cliapp.Lifecycle, error) {
	log.Info("running api...")
	cfg := config.NewConfig(ctx)
	return rest.NewAPI(ctx.Context, &cfg)
}

func NewCli(GitCommit string, GitData string) *cli.App {
	flags := flags2.Flags
	return &cli.App{
		Version:              params.VersionWithCommit(GitCommit, GitData), // 将git提交信息，和版本信息组合在一起生产版本信息
		Description:          "An exchange school services with rpc and rest api server",
		EnableBashCompletion: true,
		Commands: []*cli.Command{
			{
				Name:        "api",
				Flags:       flags,
				Description: "Run api services",
				Action:      cliapp.LifecycleCmd(runRestApi),
			},
			{
				Name:        "rpc",
				Flags:       flags,
				Description: "Run rpc services",
				Action:      cliapp.LifecycleCmd(runRpc),
			},
			{
				Name:        "migrate",
				Flags:       flags,
				Description: "Run database migrations",
				Action:      runMigrations,
			},
		},
	}
}

func runMigrations(ctx *cli.Context) error {
	ctx.Context = opio.CancelOnInterrupt(ctx.Context)
	cfg := config.NewConfig(ctx)
	db, err := database.NewDB(ctx.Context, cfg.Database)
	if err != nil {
		return err
	}
	defer func(db *database.DB) {
		err := db.Close()
		if err != nil {
		}
	}(db)
	err = db.ExecuteSQLMigration(cfg.Migrations)
	if err != nil {
		return err
	}
	return nil
}
