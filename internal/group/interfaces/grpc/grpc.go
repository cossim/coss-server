package grpc

//import (
//	"context"
//	"errors"
//	"github.com/cossim/coss-server/internal/group/domain/group"
//	"github.com/cossim/coss-server/pkg/cache"
//	"github.com/cossim/coss-server/pkg/code"
//	pkgconfig "github.com/cossim/coss-server/pkg/config"
//	"github.com/cossim/coss-server/pkg/db"
//	"github.com/cossim/coss-server/pkg/version"
//	"google.golang.org/grpc/health/grpc_health_v1"
//	"log"
//
//	//"github.com/cossim/coss-interfaces/internal/group/api/grpc/v1"
//	api "github.com/cossim/coss-server/internal/group/api/grpc/v1"
//	"github.com/cossim/coss-server/internal/group/domain/entity"
//	"github.com/cossim/coss-server/internal/group/infrastructure/persistence"
//	"google.golang.org/grpc"
//	"google.golang.org/grpc/codes"
//	"google.golang.org/grpc/health"
//	"google.golang.org/grpc/status"
//	"gorm.io/gorm"
//	"strconv"
//)
//
//type GroupServiceServer struct {
//	gr          persistence.GroupRepo
//	ac          *pkgconfig.AppConfig
//	cache       cache.GroupCache
//	cacheEnable bool
//	closeFunc   func() error
//}
//
//func (s *GroupServiceServer) Register(srv *grpc.Server) {
//	api.RegisterGroupServiceServer(srv, s)
//}
//
//func (s *GroupServiceServer) RegisterHealth(srv *grpc.Server) {
//	grpc_health_v1.RegisterHealthServer(srv, health.NewServer())
//}
//
//func (s *GroupServiceServer) Init(cfg *pkgconfig.AppConfig) error {
//	mysql, err := db.NewMySQL(cfg.MySQL.Address, strconv.Itoa(cfg.MySQL.Port), cfg.MySQL.Username, cfg.MySQL.Password, cfg.MySQL.Database, int64(cfg.Log.Level), cfg.MySQL.Opts)
//	if err != nil {
//		return err
//	}
//
//	dbConn, err := mysql.GetConnection()
//	if err != nil {
//		return err
//	}
//	infra := persistence.NewRepositories(dbConn)
//	if err = infra.Automigrate(); err != nil {
//		return err
//	}
//
//	groupCache, err := cache.NewGroupCacheRedis(cfg.Redis.Addr(), cfg.Redis.Password, 0)
//	if err != nil {
//		return err
//	}
//
//	s.closeFunc = func() error {
//		return groupCache.Close()
//	}
//
//	s.gr = infra.Gr
//	s.ac = cfg
//	s.cache = groupCache
//	s.cacheEnable = cfg.Cache.Enable
//	return nil
//}
//
//func (s *GroupServiceServer) Name() string {
//	return "group_service"
//}
//
//func (s *GroupServiceServer) Version() string {
//	return version.FullVersion()
//}
//
//func (s *GroupServiceServer) Stop(ctx context.Context) error {
//	if err := s.cache.DeleteAllCache(ctx); err != nil {
//		log.Printf("delete all group cache error: %v", err)
//	}
//	return s.closeFunc()
//}
//
//func (s *GroupServiceServer) DiscoverServices(services map[string]*grpc.ClientConn) error { return nil }
//
//func (s *GroupServiceServer) GetGroupInfoByGid(ctx context.Context, request *api.GetGroupInfoRequest) (*api.Group, error) {
//	resp := &api.Group{}
//
//	group, err := s.gr.GetGroupInfoByGid(uint(request.GetGid()))
//	if err != nil {
//		if errors.Is(err, gorm.ErrRecordNotFound) {
//			return resp, status.Error(codes.Code(code.GroupErrGroupNotFound.Code()), err.Error())
//		}
//		return resp, status.Error(codes.Code(code.GroupErrGetGroupInfoByGidFailed.Code()), err.Error())
//	}
//
//	// 将领域模型转换为 gRPC API 模型
//	resp = &api.Group{
//		Id:              uint32(group.ID),
//		Type:            api.GroupType(group.Type),
//		Status:          api.GroupStatus(group.Status),
//		MaxMembersLimit: int32(group.MaxMembersLimit),
//		CreatorId:       group.CreatorID,
//		Name:            group.Name,
//		Avatar:          group.Avatar,
//	}
//	return resp, nil
//}
//
//func (s *GroupServiceServer) GetBatchGroupInfoByIDs(ctx context.Context, request *api.GetBatchGroupInfoRequest) (*api.GetBatchGroupInfoResponse, error) {
//	resp := &api.GetBatchGroupInfoResponse{}
//
//	if len(request.GroupIds) == 0 {
//		return resp, code.MyCustomErrorCode.CustomMessage("group ids is empty")
//	}
//
//	if s.cacheEnable {
//		groups, err := s.cache.GetGroups(ctx, request.GroupIds)
//		if err == nil && len(groups) > 0 {
//			resp.Groups = groups
//			return resp, nil
//		}
//	}
//
//	//将uint32转成uint
//	groupIds := make([]uint, len(request.GroupIds))
//	for i, id := range request.GroupIds {
//		groupIds[i] = uint(id)
//	}
//
//	groups, err := s.gr.GetBatchGetGroupInfoByIDs(groupIds)
//	if err != nil {
//		return resp, status.Error(codes.Code(code.GroupErrGetBatchGroupInfoByIDsFailed.Code()), err.Error())
//	}
//
//	// 将领域模型转换为 gRPC API 模型
//	var groupAPIs []*api.Group
//	for _, group := range groups {
//		groupAPI := &api.Group{
//			Id:              uint32(group.ID),
//			Type:            api.GroupType(group.Type),
//			Status:          api.GroupStatus(group.Status),
//			MaxMembersLimit: int32(group.MaxMembersLimit),
//			CreatorId:       group.CreatorID,
//			Name:            group.Name,
//			Avatar:          group.Avatar,
//		}
//		groupAPIs = append(groupAPIs, groupAPI)
//	}
//
//	resp.Groups = groupAPIs
//
//	if s.cacheEnable {
//		if err = s.cache.SetGroup(ctx, resp.Groups...); err != nil {
//			log.Printf("set group cache failed: %v", err)
//		}
//	}
//
//	return resp, nil
//}
//
//func (s *GroupServiceServer) UpdateGroup(ctx context.Context, request *api.UpdateGroupRequest) (*api.Group, error) {
//	resp := &api.Group{}
//
//	group := &group.Group{
//		BaseModel: entity.BaseModel{
//			ID: uint(request.Group.Id),
//		},
//		Type:            group.Type(request.GetGroup().Type),
//		Status:          group.Status(request.GetGroup().Status),
//		MaxMembersLimit: int(request.GetGroup().MaxMembersLimit),
//		CreatorID:       request.GetGroup().CreatorId,
//		Name:            request.Group.Name,
//		Avatar:          request.GetGroup().Avatar,
//	}
//
//	updatedGroup, err := s.gr.Update(group)
//	if err != nil {
//		return resp, status.Error(codes.Code(code.GroupErrUpdateGroupFailed.Code()), err.Error())
//	}
//
//	// 将领域模型转换为 gRPC API 模型
//	resp = &api.Group{
//		Id:              uint32(updatedGroup.ID),
//		Type:            api.GroupType(group.Type),
//		Status:          api.GroupStatus(group.Status),
//		MaxMembersLimit: int32(updatedGroup.MaxMembersLimit),
//		CreatorId:       updatedGroup.CreatorID,
//		Name:            updatedGroup.Name,
//		Avatar:          updatedGroup.Avatar,
//	}
//
//	if s.cacheEnable {
//		if err := s.cache.SetGroup(ctx, resp); err != nil {
//			log.Printf("set group cache failed: %v", err)
//		}
//	}
//
//	return resp, nil
//}
//
//func (s *GroupServiceServer) CreateGroup(ctx context.Context, request *api.CreateGroupRequest) (*api.Group, error) {
//	resp := &api.Group{}
//
//	group := &group.Group{
//		Type:            group.Type(request.GetGroup().Type),
//		Status:          group.Status(request.GetGroup().Status),
//		MaxMembersLimit: int(request.GetGroup().MaxMembersLimit),
//		CreatorID:       request.GetGroup().CreatorId,
//		Name:            request.Group.Name,
//		Avatar:          request.GetGroup().Avatar,
//	}
//
//	createdGroup, err := s.gr.InsertGroup(group)
//	if err != nil {
//		return resp, status.Error(codes.Code(code.GroupErrInsertGroupFailed.Code()), err.Error())
//	}
//
//	// 将领域模型转换为 gRPC API 模型
//	resp = &api.Group{
//		Id:              uint32(createdGroup.ID),
//		Type:            api.GroupType(group.Type),
//		Status:          api.GroupStatus(group.Status),
//		MaxMembersLimit: int32(createdGroup.MaxMembersLimit),
//		CreatorId:       createdGroup.CreatorID,
//		Name:            createdGroup.Name,
//		Avatar:          createdGroup.Avatar,
//	}
//
//	return resp, nil
//}
//
//func (s *GroupServiceServer) CreateGroupRevert(ctx context.Context, request *api.CreateGroupRequest) (*api.Group, error) {
//	resp := &api.Group{}
//	if err := s.gr.Delete(uint(request.Group.Id)); err != nil {
//		return resp, status.Error(codes.Code(code.GroupErrDeleteGroupFailed.Code()), err.Error())
//	}
//	return resp, nil
//}
//
//func (s *GroupServiceServer) DeleteGroup(ctx context.Context, request *api.DeleteGroupRequest) (*api.EmptyResponse, error) {
//	resp := &api.EmptyResponse{}
//
//	//return resp, status.Error(codes.Aborted, errors.New("测试回滚").Error())
//
//	if err := s.gr.Delete(uint(request.GetGid())); err != nil {
//		return resp, status.Error(codes.Aborted, err.Error())
//	}
//
//	if s.cacheEnable {
//		if err := s.cache.DeleteGroup(ctx, request.Gid); err != nil {
//			log.Printf("set group cache failed: %v", err)
//		}
//	}
//
//	return resp, nil
//}
//
//func (s *GroupServiceServer) DeleteGroupRevert(ctx context.Context, request *api.DeleteGroupRequest) (*api.EmptyResponse, error) {
//	resp := &api.EmptyResponse{}
//	if err := s.gr.UpdateGroupByGroupID(uint(request.GetGid()), map[string]interfaces{}{
//		"deleted_at": 0,
//	}); err != nil {
//		return resp, status.Error(codes.Code(code.GroupErrDeleteGroupFailed.Code()), err.Error())
//	}
//	return resp, nil
//}
