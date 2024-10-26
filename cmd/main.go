package main

import (
	"context"
	"github.com/ethereum/go-ethereum/log"
	"github.com/zhulida1234/school-rpc/common/opio"
	"os"
)

var (
	GitCommit = ""
	GitData   = ""
)

func main() {
	log.SetDefault(log.NewLogger(log.NewTerminalHandlerWithLevel(os.Stderr, log.LevelInfo, true)))
	app := NewCli(GitCommit, GitData)
	// 这个方法的作用是否是增加一个信号中断处理器，用于通知给上下文
	ctx := opio.WithInterruptBlocker(context.Background())
	// 真正执行的是command.go Run方法
	if err := app.RunContext(ctx, os.Args); err != nil {
		log.Error("Application failed")
		os.Exit(1)
	}
}
