package main

import (
	"github.com/ethereum/go-ethereum/params"
	"github.com/urfave/cli/v2"
	flags2 "github.com/zhulida1234/go-rpc-service/flags"
	"github.com/zhulida1234/school-rpc/common/opio"
	"github.com/zhulida1234/school-rpc/config"
	"github.com/zhulida1234/school-rpc/database"
)

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

func NewCli(GitCommit string, GitData string) *cli.App {
	flags := flags2.Flags
	return &cli.App{
		Version:              params.VersionWithCommit(GitCommit, GitData), // 将git提交信息，和版本信息组合在一起生产版本信息
		Description:          "An exchange school services with rpc and rest api server",
		EnableBashCompletion: true,
		Commands: []*cli.Command{
			{
				Name:        "migrate",
				Flags:       flags,
				Description: "Run database migrations",
				Action:      runMigrations,
			},
		},
	}
}