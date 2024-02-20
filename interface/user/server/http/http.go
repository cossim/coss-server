package http

import (
	"context"
	"github.com/cossim/coss-server/interface/user/service"
	pkgconfig "github.com/cossim/coss-server/pkg/config"
	"github.com/cossim/coss-server/pkg/encryption"
	"github.com/cossim/coss-server/pkg/http/middleware"
	plog "github.com/cossim/coss-server/pkg/log"
	"github.com/cossim/coss-server/pkg/server"
	"github.com/cossim/coss-server/pkg/version"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

var (
	_ server.HTTPService = &Handler{}
)

type Handler struct {
	redisClient *redis.Client
	logger      *zap.Logger
	svc         *service.Service
	enc         encryption.Encryptor
	key         string
}

func (h Handler) Init(cfg *pkgconfig.AppConfig) error {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Addr(),
		Password: cfg.Redis.Password, // no password set
		DB:       0,                  // use default DB
		//Protocol: cfg,
	})
	h.redisClient = rdb
	h.logger = plog.NewDevLogger("relation_bff")
	h.enc = encryption.NewEncryptor([]byte(cfg.Encryption.Passphrase), cfg.Encryption.Name, cfg.Encryption.Email, cfg.Encryption.RsaBits, cfg.Encryption.Enable)
	h.svc = service.New(cfg)

	if err := h.enc.ReadKeyPair(); err != nil {
		return err
	}
	h.key = h.enc.GetPublicKey()
	return nil
}

func (h Handler) Name() string {
	//TODO implement me
	panic("implement me")
}

func (h Handler) Version() string {
	return version.FullVersion()
}

// @title user服务

func (h Handler) RegisterRoute(r gin.IRouter) {
	gin.SetMode(gin.ReleaseMode)
	// 添加一些中间件或其他配置
	r.Use(middleware.CORSMiddleware(), middleware.GRPCErrorMiddleware(h.logger), middleware.EncryptionMiddleware(h.enc), middleware.RecoveryMiddleware())
	u := r.Group("/api/v1/user")
	u.POST("/login", h.login)
	u.POST("/register", h.register)
	u.GET("/activate", h.userActivate)
	u.POST("/public_key/reset", h.resetUserPublicKey)
	u.POST("/email/code/send", h.sendEmailCode)
	u.GET("/system/key/get", h.GetSystemPublicKey)
	u.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.NewHandler(), ginSwagger.InstanceName("user")))

	u.Use(middleware.AuthMiddleware(h.redisClient))
	u.GET("/search", h.search)
	u.GET("/info", h.getUserInfo)
	u.POST("/logout", h.logout)
	u.POST("/info/modify", h.modifyUserInfo)
	u.POST("/password/modify", h.modifyUserPassword)
	u.POST("/key/set", h.setUserPublicKey)
	u.POST("/bundle/modify", h.modifyUserSecretBundle)
	u.GET("/bundle/get", h.getUserSecretBundle)
	u.GET("/clients/get", h.getUserLoginClients)
}

func (h Handler) Health(r gin.IRouter) string {
	//TODO implement me
	panic("implement me")
}

func (h Handler) Stop(ctx context.Context) error {
	return h.svc.Stop(ctx)
}

func (h Handler) DiscoverServices(services map[string]*grpc.ClientConn) error {
	for k, v := range services {
		if err := h.svc.HandlerGrpcClient(k, v); err != nil {
			h.logger.Error("handler grpc client error", zap.String("name", k), zap.String("addr", v.Target()), zap.Error(err))
		}
	}
	return nil
}
