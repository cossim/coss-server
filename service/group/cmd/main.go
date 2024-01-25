package main

import (
	"flag"
	"fmt"
	"github.com/cossim/coss-server/pkg/db"
	api "github.com/cossim/coss-server/service/group/api/v1"
	"github.com/cossim/coss-server/service/group/config"
	"github.com/cossim/coss-server/service/group/infrastructure/persistence"
	"github.com/cossim/coss-server/service/group/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

var discover bool

func init() {
	flag.StringVar(&config.ConfigFile, "config", "/config/config.yaml", "Path to configuration file")
	flag.BoolVar(&discover, "discover", false, "Enable service discovery")
	flag.Parse()
}

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
	svc := service.NewService(infra, config.Conf)
	api.RegisterGroupServiceServer(grpcServer, svc)
	// 注册服务开启健康检查
	grpc_health_v1.RegisterHealthServer(grpcServer, health.NewServer())
	log.Printf("gRPC server is running on addr: %s\n", config.Conf.GRPC.Addr())

	go func() {
		if err = grpcServer.Serve(lis); err != nil {
			panic(err)
		}
	}()

	svc.Start(discover)

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			grpcServer.Stop()
			svc.Close(discover)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
