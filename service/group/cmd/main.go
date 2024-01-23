package main

import (
	"fmt"
	"github.com/cossim/coss-server/pkg/db"
	api "github.com/cossim/coss-server/service/group/api/v1"
	"github.com/cossim/coss-server/service/group/config"
	"github.com/cossim/coss-server/service/group/infrastructure/persistence"
	"github.com/cossim/coss-server/service/group/service"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	if err := config.Init(); err != nil {
		panic(err)
	}

	// 启动 gRPC 服务器
	lis, err := net.Listen("tcp", fmt.Sprintf("%s", config.Conf.GRPC.Addr()))
	if err != nil {
		panic(err)
	}

	dbConn, err := db.NewMySQLFromDSN(config.Conf.MySQL.DSN).GetConnection()
	if err != nil {
		panic(err)
	}

	infra := persistence.NewRepositories(dbConn)
	if err = infra.Automigrate(); err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	svc := service.NewService(infra)
	api.RegisterGroupServiceServer(grpcServer, svc)

	fmt.Printf("gRPC server is running on addr: %s\n", config.Conf.GRPC.Addr())

	go func() {
		if err = grpcServer.Serve(lis); err != nil {
			panic(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			grpcServer.Stop()
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
