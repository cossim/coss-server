package http

import (
	"context"
	"github.com/cossim/coss-server/internal/push/service"
	"github.com/cossim/coss-server/internal/user/cache"
	pkgconfig "github.com/cossim/coss-server/pkg/config"
	"github.com/cossim/coss-server/pkg/encryption"
	"github.com/cossim/coss-server/pkg/http/middleware"
	plog "github.com/cossim/coss-server/pkg/log"
	"github.com/cossim/coss-server/pkg/version"
	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"log"
)

type Handler struct {
	logger       *zap.Logger
	enc          encryption.Encryptor
	PushService  *service.Service
	userCache    cache.UserCache
	socketServer *socketio.Server
	jwtSecret    string
}

func (h *Handler) Init(cfg *pkgconfig.AppConfig) error {
	h.logger = plog.NewDefaultLogger("push_bff", int8(cfg.Log.Level))
	if cfg.Encryption.Enable {
		return h.enc.ReadKeyPair()
	}
	userCache, err := cache.NewUserCacheRedis(cfg.Redis.Addr(), cfg.Redis.Password, 0)
	if err != nil {
		return err
	}
	h.userCache = userCache
	h.socketServer = h.PushService.SocketServer
	h.setupSocketEvent()
	h.jwtSecret = cfg.SystemConfig.JwtSecret
	h.enc = encryption.NewEncryptor([]byte(cfg.Encryption.Passphrase), cfg.Encryption.Name, cfg.Encryption.Email, cfg.Encryption.RsaBits, cfg.Encryption.Enable)
	return nil
}

func (h *Handler) Name() string {
	return "push_bff"
}

func (h *Handler) Version() string {
	return version.FullVersion()
}

func (h *Handler) RegisterRoute(r gin.IRouter) {
	u := r.Group("/api/v1/push")
	u.Use(middleware.CORSMiddleware(), middleware.GRPCErrorMiddleware(h.logger), middleware.EncryptionMiddleware(h.enc), middleware.RecoveryMiddleware())
	u.Use(middleware.AuthMiddleware(h.userCache, h.jwtSecret))
	u.GET("/ws/*any", gin.WrapH(h.socketServer))
	u.POST("/ws/*any", gin.WrapH(h.socketServer))

}

func (h *Handler) Health(r gin.IRouter) string {
	//TODO implement me
	panic("implement me")
}

func (h *Handler) Stop(ctx context.Context) error {
	err := h.socketServer.Close()
	if err != nil {
		return err
	}
	return nil
}

func (h *Handler) DiscoverServices(services map[string]*grpc.ClientConn) error {
	for k, v := range services {
		if err := h.PushService.HandlerGrpcClient(k, v); err != nil {
			h.logger.Error("handler grpc client error", zap.String("name", k), zap.String("addr", v.Target()), zap.Error(err))
		}
	}
	return nil
}

func (h *Handler) setupSocketEvent() {
	if h.socketServer == nil {
		panic("socketio server is nil")
	}
	h.socketServer.OnConnect("/", h.ws)
	h.socketServer.OnError("/", h.error)
	h.socketServer.OnDisconnect("/", h.disconnect)
	h.socketServer.OnEvent("/", "reply", h.reply)
	h.socketServer.OnEvent("/", "bye", h.bye)

	go func() {
		if err := h.socketServer.Serve(); err != nil {
			log.Fatalf("socketio listen error: %s\n", err)
		}
	}()
}
