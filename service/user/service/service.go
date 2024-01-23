package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/cossim/coss-server/pkg/code"
	"github.com/cossim/coss-server/pkg/utils/time"
	"github.com/cossim/coss-server/service/user/domain/entity"
	"github.com/cossim/coss-server/service/user/domain/repository"
	"github.com/cossim/coss-server/service/user/utils"
	codes "google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"

	api "github.com/cossim/coss-server/service/user/api/v1"
)

func NewService(ur repository.UserRepository) *Service {
	return &Service{
		ur: ur,
	}
}

type Service struct {
	ur repository.UserRepository
	api.UnimplementedUserServiceServer
}

// 用户登录
func (g *Service) UserLogin(ctx context.Context, request *api.UserLoginRequest) (*api.UserLoginResponse, error) {
	resp := &api.UserLoginResponse{}

	userInfo, err := g.ur.GetUserInfoByEmail(request.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return resp, status.Error(codes.Code(code.UserErrNotExistOrPassword.Code()), err.Error())
		}
		return resp, status.Error(codes.Code(code.UserErrLoginFailed.Code()), err.Error())
	}

	if userInfo.Password != request.Password {
		return resp, status.Error(codes.Code(code.UserErrNotExistOrPassword.Code()), code.UserErrNotExistOrPassword.Message())
	}
	//修改登录时间
	userInfo.LastAt = time.Now()
	_, err = g.ur.UpdateUser(userInfo)
	if err != nil {
		return resp, status.Error(codes.Code(code.UserErrLoginFailed.Code()), err.Error())
	}

	switch userInfo.Status {
	case entity.UserStatusLock:
		return resp, status.Error(codes.Code(code.UserErrLocked.Code()), code.UserErrLocked.Message())
	case entity.UserStatusDeleted:
		return resp, status.Error(codes.Code(code.UserErrDeleted.Code()), code.UserErrDeleted.Message())
	case entity.UserStatusDisable:
		return resp, status.Error(codes.Code(code.UserErrDisabled.Code()), code.UserErrDisabled.Message())
	case entity.UserStatusNormal:
		return &api.UserLoginResponse{
			UserId:    userInfo.ID,
			NickName:  userInfo.NickName,
			Email:     userInfo.Email,
			Tel:       userInfo.Tel,
			Avatar:    userInfo.Avatar,
			Signature: userInfo.Signature,
		}, nil
	default:
		return resp, status.Error(codes.Code(code.UserErrStatusException.Code()), code.UserErrStatusException.Message())
	}
}

// 用户注册
func (g *Service) UserRegister(ctx context.Context, request *api.UserRegisterRequest) (*api.UserRegisterResponse, error) {
	resp := &api.UserRegisterResponse{}
	//添加用户
	_, err := g.ur.GetUserInfoByEmail(request.Email)
	if err == nil {
		return resp, status.Error(codes.Code(code.UserErrEmailAlreadyRegistered.Code()), code.UserErrEmailAlreadyRegistered.Message())
	}
	userInfo, err := g.ur.InsertUser(&entity.User{
		Email:     request.Email,
		Password:  request.Password,
		NickName:  request.NickName,
		Avatar:    request.Avatar,
		PublicKey: request.PublicKey,
		//Action:   entity.UserStatusLock,
		Status: entity.UserStatusNormal,
		ID:     utils.GenUUid(),
	})
	if err != nil {
		return resp, status.Error(codes.Code(code.UserErrRegistrationFailed.Code()), err.Error())
	}
	resp.UserId = userInfo.ID
	return resp, nil
}

func (g *Service) UserInfo(ctx context.Context, request *api.UserInfoRequest) (*api.UserInfoResponse, error) {
	resp := &api.UserInfoResponse{}
	userInfo, err := g.ur.GetUserInfoByUid(request.UserId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return resp, status.Error(codes.Code(code.UserErrNotExistOrPassword.Code()), err.Error())
		}
		return resp, status.Error(codes.Code(code.UserErrGetUserInfoFailed.Code()), err.Error())
	}
	resp = &api.UserInfoResponse{
		UserId:    userInfo.ID,
		NickName:  userInfo.NickName,
		Email:     userInfo.Email,
		Tel:       userInfo.Tel,
		Avatar:    userInfo.Avatar,
		Signature: userInfo.Signature,
		Status:    api.UserStatus(userInfo.Status),
	}
	return resp, nil
}

func (g *Service) GetBatchUserInfo(ctx context.Context, request *api.GetBatchUserInfoRequest) (*api.GetBatchUserInfoResponse, error) {
	resp := &api.GetBatchUserInfoResponse{}
	users, err := g.ur.GetBatchGetUserInfoByIDs(request.UserIds)
	if err != nil {
		fmt.Printf("无法获取用户列表信息: %v\n", err)
		return nil, status.Error(codes.Code(code.UserErrUnableToGetUserListInfo.Code()), err.Error())
	}
	for _, user := range users {
		resp.Users = append(resp.Users, &api.UserInfoResponse{
			UserId:    user.ID,
			NickName:  user.NickName,
			Email:     user.Email,
			Tel:       user.Tel,
			Avatar:    user.Avatar,
			Signature: user.Signature,
			Status:    api.UserStatus(user.Status),
		})
	}

	return resp, nil
}

func (g *Service) GetUserInfoByEmail(ctx context.Context, request *api.GetUserInfoByEmailRequest) (*api.UserInfoResponse, error) {
	resp := &api.UserInfoResponse{}
	userInfo, err := g.ur.GetUserInfoByEmail(request.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return resp, status.Error(codes.Code(code.UserErrNotExistOrPassword.Code()), err.Error())
		}
		return resp, status.Error(codes.Code(code.UserErrGetUserInfoFailed.Code()), err.Error())
	}
	resp = &api.UserInfoResponse{
		UserId:    userInfo.ID,
		NickName:  userInfo.NickName,
		Email:     userInfo.Email,
		Tel:       userInfo.Tel,
		Avatar:    userInfo.Avatar,
		Signature: userInfo.Signature,
		Status:    api.UserStatus(userInfo.Status),
	}
	return resp, nil
}

func (g *Service) GetUserPublicKey(ctx context.Context, in *api.UserRequest) (*api.GetUserPublicKeyResponse, error) {
	key, err := g.ur.GetUserPublicKey(in.UserId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &api.GetUserPublicKeyResponse{}, status.Error(codes.Code(code.UserErrPublicKeyNotExist.Code()), err.Error())
		}
		return &api.GetUserPublicKeyResponse{}, status.Error(codes.Code(code.UserErrGetUserPublicKeyFailed.Code()), err.Error())
	}
	return &api.GetUserPublicKeyResponse{PublicKey: key}, nil
}

func (g *Service) SetUserPublicKey(ctx context.Context, in *api.SetPublicKeyRequest) (*api.UserResponse, error) {
	if err := g.ur.SetUserPublicKey(in.UserId, in.PublicKey); err != nil {
		return &api.UserResponse{}, status.Error(codes.Code(code.UserErrSaveUserPublicKeyFailed.Code()), err.Error())
	}
	return &api.UserResponse{UserId: in.UserId}, nil
}

func (g *Service) ModifyUserInfo(ctx context.Context, in *api.User) (*api.UserResponse, error) {
	resp := &api.UserResponse{}
	user, err := g.ur.UpdateUser(&entity.User{
		ID:        in.UserId,
		NickName:  in.NickName,
		Avatar:    in.Avatar,
		Signature: in.Signature,
		Status:    entity.UserStatus(in.Status),
		Tel:       in.Tel,
	})
	if err != nil {
		return resp, err
	}
	resp = &api.UserResponse{UserId: user.ID}
	return resp, nil
}

func (g *Service) ModifyUserPassword(ctx context.Context, in *api.ModifyUserPasswordRequest) (*api.UserResponse, error) {
	resp := &api.UserResponse{}
	user, err := g.ur.UpdateUser(&entity.User{
		ID:       in.UserId,
		Password: in.Password,
	})
	if err != nil {
		return resp, err
	}
	resp = &api.UserResponse{UserId: user.ID}
	return resp, nil
}

func (g *Service) GetUserPasswordByUserId(ctx context.Context, in *api.UserRequest) (*api.GetUserPasswordByUserIdResponse, error) {
	resp := &api.GetUserPasswordByUserIdResponse{}
	userInfo, err := g.ur.GetUserInfoByUid(in.UserId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return resp, status.Error(codes.Code(code.UserErrNotExistOrPassword.Code()), err.Error())
		}
		return resp, status.Error(codes.Code(code.UserErrGetUserInfoFailed.Code()), err.Error())
	}
	resp = &api.GetUserPasswordByUserIdResponse{
		Password: userInfo.Password,
		UserId:   userInfo.ID,
	}
	return resp, nil
}

func (g *Service) SetUserSecretBundle(ctx context.Context, in *api.SetUserSecretBundleRequest) (*api.SetUserSecretBundleResponse, error) {
	var resp = &api.SetUserSecretBundleResponse{}
	if err := g.ur.SetUserSecretBundle(in.UserId, in.SecretBundle); err != nil {
		return resp, status.Error(codes.Code(code.UserErrSetUserSecretBundleFailed.Code()), err.Error())
	}
	return resp, nil
}

func (g *Service) GetUserSecretBundle(ctx context.Context, in *api.GetUserSecretBundleRequest) (*api.GetUserSecretBundleResponse, error) {
	var resp = &api.GetUserSecretBundleResponse{}
	secretBundle, err := g.ur.GetUserSecretBundle(in.UserId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return resp, status.Error(codes.Code(code.UserErrNotExist.Code()), err.Error())
		}
		return resp, status.Error(codes.Code(code.UserErrGetUserSecretBundleFailed.Code()), err.Error())
	}
	resp.SecretBundle = secretBundle
	resp.UserId = in.UserId
	return resp, nil
}
