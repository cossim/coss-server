package service

import (
	storagev1 "github.com/cossim/coss-server/internal/storage/api/grpc/v1"
	grpchandler "github.com/cossim/coss-server/internal/storage/interface/grpc"
	usergrpcv1 "github.com/cossim/coss-server/internal/user/api/grpc/v1"
	pkgconfig "github.com/cossim/coss-server/pkg/config"
	plog "github.com/cossim/coss-server/pkg/log"
	"github.com/cossim/coss-server/pkg/storage"
	"github.com/cossim/coss-server/pkg/storage/minio"
	"github.com/rs/xid"
	"go.uber.org/zap"
)

type Service struct {
	ac     *pkgconfig.AppConfig
	logger *zap.Logger

	userService    usergrpcv1.UserServiceClient
	sp             storage.StorageProvider
	storageService storagev1.StorageServiceServer

	sid            string
	downloadURL    string
	gatewayAddress string
	gatewayPort    string
	dis            bool
	cache          bool
}

func New(ac *pkgconfig.AppConfig, grpcService *grpchandler.Handler) *Service {
	logger := setupLogger(ac)
	svc := &Service{
		downloadURL: "/api/v1/storage/files/download",
		sp:          setMinIOProvider(ac),
		logger:      logger,
		ac:          ac,
		sid:         xid.New().String(),
		dis:         false,
	}
	svc.setLoadSystem()
	svc.storageService = grpcService
	return svc
}

func setMinIOProvider(ac *pkgconfig.AppConfig) storage.StorageProvider {
	var err error
	sp, err := minio.NewMinIOStorage(ac.OSS.Addr(), ac.OSS.AccessKey, ac.OSS.SecretKey, ac.OSS.SSL)
	if err != nil {
		panic(err)
	}

	return sp
}

func setupLogger(c *pkgconfig.AppConfig) *zap.Logger {
	return plog.NewDefaultLogger("group_bff", int8(c.Log.Level))
}

func (s *Service) setLoadSystem() {
	s.gatewayAddress = s.ac.SystemConfig.GatewayAddress
	s.gatewayPort = s.ac.SystemConfig.GatewayPort
	if !s.ac.SystemConfig.Ssl {
		s.gatewayAddress = s.gatewayAddress + ":" + s.gatewayPort
	}
}
