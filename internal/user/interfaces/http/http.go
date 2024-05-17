package http

//
//import (
//	"context"
//	authv1 "github.com/cossim/coss-server/internal/user/api/grpc/v1"
//	grpchandler "github.com/cossim/coss-server/internal/user/interfaces/grpc"
//	"github.com/cossim/coss-server/internal/user/rpc/client"
//	"github.com/cossim/coss-server/internal/user/service"
//	pkgconfig "github.com/cossim/coss-server/pkg/config"
//	"github.com/cossim/coss-server/pkg/encryption"
//	"github.com/cossim/coss-server/pkg/http/middleware"
//	plog "github.com/cossim/coss-server/pkg/log"
//	"github.com/cossim/coss-server/pkg/manager/server"
//	"github.com/cossim/coss-server/pkg/version"
//	"github.com/gin-gonic/gin"
//	"go.uber.org/zap"
//	"google.golang.org/grpc"
//)
//
//var (
//	_ server.HTTPService = &Handler{}
//)
//
//type Handler struct {
//	logger      *zap.Logger
//	svc         *service.Service
//	enc         encryption.Encryptor
//	key         string
//	UserClient  *grpchandler.UserServiceServer
//	authService authv1.UserAuthServiceClient
//}
//
//func (h *Handler) Init(cfg *pkgconfig.AppConfig) error {
//	h.logger = plog.NewDefaultLogger("user_bff", int8(cfg.Log.Level))
//	if cfg.Encryption.Enable {
//		if err := h.enc.ReadKeyPair(); err != nil {
//			return err
//		}
//	}
//	h.enc = encryption.NewEncryptor([]byte(cfg.Encryption.Passphrase), cfg.Encryption.Name, cfg.Encryption.Email, cfg.Encryption.RsaBits, cfg.Encryption.Enable)
//	h.svc = service.New(cfg, h.UserClient)
//	var userAddr = cfg.GRPC.Addr()
//	//if cfg.Discovers["user"].Direct {
//	//	userAddr = cfg.Discovers["user"].Addr()
//	//} else {
//	//	userAddr = discovery.GetBalanceAddr(cfg.Register.Addr(), cfg.Discovers["user"].Name)
//	//}
//	authClient, err := client.NewAuthClient(userAddr)
//	if err != nil {
//		return err
//	}
//	h.authService = authClient
//	return nil
//}
//
//func (h *Handler) Name() string {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (h *Handler) Version() string {
//	return version.FullVersion()
//}
//
//// @title CossApi
//
//func (h *Handler) RegisterRoute(r gin.IRouter) {
//	gin.SetMode(gin.ReleaseMode)
//	// 添加一些中间件或其他配置
//	r.Use(middleware.CORSMiddleware(), middleware.GRPCErrorMiddleware(h.logger), middleware.EncryptionMiddleware(h.enc), middleware.RecoveryMiddleware())
//	u := r.Group("/api/v1/user")
//	u.POST("/login", h.login)
//	u.POST("/register", h.register)
//	u.GET("/activate", h.userActivate)
//	u.POST("/email/code/send", h.sendEmailCode)
//	u.GET("/system/key/get", h.getSystemPublicKey)
//
//	u.Use(middleware.AuthMiddleware(h.authService))
//	u.POST("/public_key/reset", h.resetUserPublicKey)
//	u.GET("/search", h.search)
//	u.GET("/info", h.getUserInfo)
//	u.POST("/logout", h.logout)
//	u.POST("/info/modify", h.modifyUserInfo)
//	u.POST("/password/modify", h.modifyUserPassword)
//	u.POST("/key/set", h.setUserPublicKey)
//	u.POST("/bundle/modify", h.modifyUserSecretBundle)
//	u.GET("/bundle/get", h.getUserSecretBundle)
//	u.GET("/clients/get", h.getUserLoginClients)
//	//修改用户头像
//	u.POST("/avatar/modify", h.modifyUserAvatar)
//}
//
//func (h *Handler) Health(r gin.IRouter) string {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (h *Handler) Stop(ctx context.Context) error {
//	return h.svc.Stop(ctx)
//}
//
//func (h *Handler) DiscoverServices(services map[string]*grpc.ClientConn) error {
//	for k, v := range services {
//		if err := h.svc.HandlerGrpcClient(k, v); err != nil {
//			h.logger.Error("handler grpc client error", zap.String("name", k), zap.String("addr", v.Target()), zap.Error(err))
//		}
//	}
//	return nil
//}
