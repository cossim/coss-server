package service

import (
	"context"
	"fmt"
	relationgrpcv1 "github.com/cossim/coss-server/internal/relation/api/grpc/v1"
	grpchandler "github.com/cossim/coss-server/internal/user/interface/grpc"
	"github.com/cossim/coss-server/pkg/cache"
	pkgconfig "github.com/cossim/coss-server/pkg/config"
	"github.com/cossim/coss-server/pkg/email"
	"github.com/cossim/coss-server/pkg/email/smtp"
	plog "github.com/cossim/coss-server/pkg/log"
	"github.com/cossim/coss-server/pkg/msg_queue"
	"github.com/cossim/coss-server/pkg/storage"
	"github.com/cossim/coss-server/pkg/storage/minio"
	"github.com/rs/xid"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"time"
)

// Service struct
type Service struct {
	ac         *pkgconfig.AppConfig
	logger     *zap.Logger
	userClient *grpchandler.Handler
	//userLoginClient usergrpc.UserLoginServiceClient
	relClient      relationgrpcv1.UserRelationServiceClient
	dialogClient   relationgrpcv1.DialogServiceClient
	sp             storage.StorageProvider
	redisClient    *cache.RedisClient
	smtpClient     email.EmailProvider
	rabbitMQClient *msg_queue.RabbitMQ
	dtmGrpcServer  string
	userGrpcServer string

	sid             string
	tokenExpiration time.Duration
	downloadURL     string
	gatewayAddress  string
	gatewayPort     string
	cache           bool
}

func New(ac *pkgconfig.AppConfig, grpcService *grpchandler.Handler) (s *Service) {
	s = &Service{
		ac:              ac,
		logger:          plog.NewDefaultLogger("user_bff", int8(ac.Log.Level)),
		sid:             xid.New().String(),
		tokenExpiration: 60 * 60 * 24 * 3 * time.Second,
		rabbitMQClient:  setRabbitMQProvider(ac),
		//appPath:         path,
		sp:            setMinIOProvider(ac),
		dtmGrpcServer: ac.Dtm.Addr(),
		downloadURL:   "/api/v1/storage/files/download",
		smtpClient:    setupSmtpProvider(ac),
	}
	s.setLoadSystem()
	s.setupRedisClient()
	s.cache = s.setCacheConfig()
	s.userClient = grpcService
	return s
}

func (s *Service) HandlerGrpcClient(serviceName string, conn *grpc.ClientConn) error {
	switch serviceName {
	case "relation_service":
		s.relClient = relationgrpcv1.NewUserRelationServiceClient(conn)
		s.dialogClient = relationgrpcv1.NewDialogServiceClient(conn)
		s.logger.Info("gRPC client for relation service initialized", zap.String("addr", conn.Target()))
	}

	return nil
}

func (s *Service) setCacheConfig() bool {
	if s.redisClient == nil && s.ac.Cache.Enable {
		panic("redis is nil")
	}
	return s.ac.Cache.Enable
}

func (s *Service) setupRedisClient() {
	s.redisClient = cache.NewRedisClient(s.ac.Redis.Addr(), s.ac.Redis.Password)
}

func (s *Service) Ping() {
}

func setRabbitMQProvider(ac *pkgconfig.AppConfig) *msg_queue.RabbitMQ {
	rmq, err := msg_queue.NewRabbitMQ(fmt.Sprintf("amqp://%s:%s@%s", ac.MessageQueue.Username, ac.MessageQueue.Password, ac.MessageQueue.Addr()))
	if err != nil {
		panic(err)
	}
	return rmq
}

func setMinIOProvider(ac *pkgconfig.AppConfig) storage.StorageProvider {
	var err error
	fmt.Println("ac.OSS => ", ac.OSS)
	sp, err := minio.NewMinIOStorage(ac.OSS.Addr(), ac.OSS.AccessKey, ac.OSS.SecretKey, ac.OSS.SSL)
	if err != nil {
		panic(err)
	}

	return sp
}

func setupSmtpProvider(ac *pkgconfig.AppConfig) email.EmailProvider {
	smtpStorage, err := smtp.NewSmtpStorage(ac.Email.SmtpServer, ac.Email.Port, ac.Email.Username, ac.Email.Password)
	if err != nil {
		panic(err)
	}
	return smtpStorage
}

func (s *Service) setLoadSystem() {

	env := s.ac.SystemConfig.Environment
	if env == "" {
		env = "dev"
	}

	switch env {
	case "prod":
		gatewayAdd := s.ac.SystemConfig.GatewayAddress
		if gatewayAdd == "" {
			gatewayAdd = "43.229.28.107"
		}

		s.gatewayAddress = gatewayAdd

		gatewayPo := s.ac.SystemConfig.GatewayPort
		if gatewayPo == "" {
			gatewayPo = "8080"
		}
		s.gatewayPort = gatewayPo
	default:
		gatewayAdd := s.ac.SystemConfig.GatewayAddressDev
		if gatewayAdd == "" {
			gatewayAdd = "127.0.0.1"
		}

		s.gatewayAddress = gatewayAdd

		gatewayPo := s.ac.SystemConfig.GatewayPortDev
		if gatewayPo == "" {
			gatewayPo = "8080"
		}
		s.gatewayPort = gatewayPo
	}
	if !s.ac.SystemConfig.Ssl {
		s.gatewayAddress = s.gatewayAddress + ":" + s.gatewayPort
	}
}

func (s *Service) Stop(ctx context.Context) error {
	return nil
}
