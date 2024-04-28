package grpc

import (
	"context"
	"fmt"
	api "github.com/cossim/coss-server/internal/relation/api/grpc/v1"
	"github.com/cossim/coss-server/internal/relation/infra/persistence"
	pkgconfig "github.com/cossim/coss-server/pkg/config"
	"github.com/cossim/coss-server/pkg/db"
	"github.com/cossim/coss-server/pkg/version"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"path/filepath"
	"runtime"
	"strconv"
)

type RelationServiceServer struct {
	ac                             *pkgconfig.AppConfig
	UserServiceServer              *userServiceServer
	GroupServiceServer             *groupServiceServer
	DialogServiceServer            *dialogServiceServer
	GroupAnnouncementServer        *groupAnnouncementServer
	GroupJoinRequestServiceServer  *groupJoinRequestServiceServer
	UserFriendRequestServiceServer *userFriendRequestServiceServer
}

func (s *RelationServiceServer) Init(cfg *pkgconfig.AppConfig) error {
	mysql, err := db.NewMySQL(cfg.MySQL.Address, strconv.Itoa(cfg.MySQL.Port), cfg.MySQL.Username, cfg.MySQL.Password, cfg.MySQL.Database, int64(cfg.Log.Level), cfg.MySQL.Opts)
	if err != nil {
		return err
	}

	//userCache, err := cache.NewRelationUserCacheRedis(cfg.Redis.Addr(), cfg.Redis.Password, 0)
	//if err != nil {
	//	return err
	//}
	//
	//groupCache, err := cache.NewRelationGroupCacheRedis(cfg.Redis.Addr(), cfg.Redis.Password, 0)
	//if err != nil {
	//	return err
	//}

	dbConn, err := mysql.GetConnection()
	if err != nil {
		return err
	}

	infra := persistence.NewRepositories(dbConn)
	if err = infra.Automigrate(); err != nil {
		return err
	}

	s.ac = cfg
	s.UserServiceServer = &userServiceServer{
		db:  dbConn,
		urr: infra.Urr,
		dr:  infra.Dr,
	}
	s.GroupServiceServer = &groupServiceServer{
		db:  dbConn,
		grr: infra.Grr,
		dr:  infra.Dr,
	}
	s.DialogServiceServer = &dialogServiceServer{
		db:  dbConn,
		dr:  infra.Dr,
		dur: infra.Dur,
	}
	s.GroupJoinRequestServiceServer = &groupJoinRequestServiceServer{
		db:   dbConn,
		dr:   infra.Dr,
		grr:  infra.Grr,
		gjqr: infra.Gjqr,
	}
	s.UserFriendRequestServiceServer = &userFriendRequestServiceServer{
		db:   dbConn,
		ufqr: infra.Ufqr,
	}
	s.GroupAnnouncementServer = &groupAnnouncementServer{
		db:  dbConn,
		gar: infra.GAr,
	}
	return nil
}

func (s *RelationServiceServer) Name() string {
	//TODO implement me
	return "relation_service"
}

func (s *RelationServiceServer) Version() string { return version.FullVersion() }

func (s *RelationServiceServer) Register(srv *grpc.Server) {
	api.RegisterUserRelationServiceServer(srv, s.UserServiceServer)
	api.RegisterGroupRelationServiceServer(srv, s.GroupServiceServer)
	api.RegisterDialogServiceServer(srv, s.DialogServiceServer)
	api.RegisterUserFriendRequestServiceServer(srv, s.UserFriendRequestServiceServer)
	api.RegisterGroupJoinRequestServiceServer(srv, s.GroupJoinRequestServiceServer)
	api.RegisterGroupAnnouncementServiceServer(srv, s.GroupAnnouncementServer)
}

func (s *RelationServiceServer) RegisterHealth(srv *grpc.Server) {
	grpc_health_v1.RegisterHealthServer(srv, health.NewServer())
}

func (s *RelationServiceServer) Stop(ctx context.Context) error {
	//return s.UserServiceServer.cache.DeleteAllCache(ctx)
	return nil
}

func (s *RelationServiceServer) DiscoverServices(services map[string]*grpc.ClientConn) error {
	//TODO implement me
	return nil
}

func getFunctionName() string {
	pc, _, _, _ := runtime.Caller(1)
	return runtime.FuncForPC(pc).Name()
}

func formatErrorMessage(err error) string {
	funcName := getFunctionName()
	_, file := filepath.Split(funcName)
	return fmt.Sprintf("[%s] %s: %v", file, funcName, err)
}
