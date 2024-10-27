package services

import (
	"context"
	"fmt"
	"github.com/zhulida1234/school-rpc/database"
	"github.com/zhulida1234/school-rpc/protobuf/clazz"
	"github.com/zhulida1234/school-rpc/protobuf/student"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"sync/atomic"
)

const MaxRecvMessageSize = 1024 * 1024 * 300

type RpcServerConfig struct {
	GrpcHost string
	GrpcPort string
}

type RpcServer struct {
	*RpcServerConfig
	db     *database.DB
	stoped atomic.Bool
}

func (s *RpcServer) Stop(ctx context.Context) error {
	s.stoped.Store(true)
	return nil
}

func (s *RpcServer) Stopped() bool {
	return false
}

func (s *RpcServer) GetRpcSchoolDB() *database.SchoolDB {
	return s.db.SchoolDB
}

func NewRpcServer(config *RpcServerConfig, db *database.DB) (*RpcServer, error) {
	return &RpcServer{
		RpcServerConfig: config,
		db:              db,
	}, nil
}

func (s *RpcServer) Start(ctx context.Context) error {
	go func(s *RpcServer) {
		addr := fmt.Sprintf("%s:%s", s.GrpcHost, s.GrpcPort)
		fmt.Println("start rpc server", "addr", addr)
		listener, err := net.Listen("tcp", addr)
		if err != nil {
			fmt.Println("Could not start rpc server", "err", err)
		}

		opt := grpc.MaxRecvMsgSize(MaxRecvMessageSize)
		//创建一个新的 gRPC 服务器实例 gs，并注册反射服务（允许客户端通过反射查询服务信息）。
		gs := grpc.NewServer(opt, grpc.ChainUnaryInterceptor(nil))
		reflection.Register(gs)
		//注册服务 学生班级
		student.RegisterStudentServiceServer(gs, s)
		clazz.RegisterClazzServiceServer(gs, s)
		// 启动grpc服务
		fmt.Println("start rpc server", "port", s.GrpcPort, "address", listener.Addr())
		if err := gs.Serve(listener); err != nil {
			fmt.Println("start rpc server", "err", err)
		}
	}(s)
	return nil
}
