package auth

import (
	"errors"
	"fmt"
	"github.com/cossim/coss-server/pkg/cache"
	"github.com/cossim/coss-server/pkg/utils"
	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Claims struct {
	UserId string `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

func NewAuthenticator(db *gorm.DB, rdb *redis.Client) *Authenticator {
	return &Authenticator{db, rdb}
}

type Authenticator struct {
	DB  *gorm.DB
	RDB *redis.Client
}

const _queryUser = "SELECT * FROM users WHERE id = ?"

func (a *Authenticator) ValidateToken(tokenString string, driverType string) (bool, error) {
	token, claims, err := utils.ParseToken(tokenString)
	if err != nil || !token.Valid {
		return false, fmt.Errorf("token validation failed: %s", err.Error())
	}
	type User struct {
		ID     string `json:"id"`
		Status int64  `json:"status"`
	}
	//data, err := cache.GetAllListValues(a.RDB, claims.UserId)
	//if err != nil {
	//	fmt.Println("error => ", err)
	//	return false, err
	//}
	keys, err := cache.ScanKeys(a.RDB, claims.UserId+":"+driverType+":*")
	if err != nil {
		fmt.Println("error => ", err)
		return false, err
	}
	if len(keys) <= 0 {
		return false, errors.New("token not found")
	}

	var found = false

	for _, key := range keys {
		v, err := cache.GetKey(a.RDB, key)
		if err != nil {
			return false, err
		}
		data := v.(string)
		info, err := cache.GetUserInfo(data)
		if err != nil {
			return false, err
		}
		if info.Token == tokenString {
			found = true
		}
	}

	//users, err := cache.GetUserInfoList(data)
	//if err != nil {
	//	return false, err
	//}
	//for _, user := range users {
	//	if user.Token == tokenString {
	//	}
	//}
	if !found {
		return false, errors.New("token not found")
	}
	var user User
	if err = a.DB.Raw(_queryUser, claims.UserId).Scan(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, errors.New("user not found")
		}
		return false, fmt.Errorf("error retrieving user: %s", err.Error())
	}

	if user.Status != 1 {
		return false, errors.New("user status is abnormal")
	}

	return true, nil
}
