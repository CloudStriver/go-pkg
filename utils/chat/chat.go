package chat

import (
	"context"
	"fmt"
	"github.com/carlmjohnson/requests"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

const UserPlatformID = 5
const AdminPlatformID = 10
const ChatToken = "ChatToken"
const ChatApiUrl = "172.28.0.9:10002"

type GetTokenReq struct {
	Secret     string `json:"secret"`
	PlatformID int64  `json:"platformID"`
	UserID     string `json:"userID"`
}

type GetTokenResp struct {
	ErrCode int    `json:"errCode"`
	ErrMsg  string `json:"errMsg"`
	ErrDlt  string `json:"errDlt"`
	Data    struct {
		Token             string `json:"token"`
		ExpireTimeSeconds int    `json:"expireTimeSeconds"`
	} `json:"data"`
}

func GetToken(Redis *redis.Redis, UserID, Secret, AdminUserID string) (string, error) {
	GetTokenReq := &GetTokenReq{
		Secret:     Secret,
		PlatformID: UserPlatformID,
		UserID:     UserID,
	}
	if UserID == AdminUserID {
		GetTokenReq.PlatformID = AdminPlatformID
	}
	var GetTokenResp GetTokenResp

	if UserID == AdminUserID {
		Token, err := Redis.Get(fmt.Sprintf("%s:%s", ChatToken, UserID))
		if err != nil {
			logx.Errorf("Redis获取缓存异常[%v]\n", err)
		}
		if Token != "" {
			return Token, nil
		} else {

			err := requests.URL(fmt.Sprintf("http://%s/auth/user_token", ChatApiUrl)).BodyJSON(GetTokenReq).Header("operationID", "1").ToJSON(&GetTokenResp).Fetch(context.Background())
			if err != nil {
				logx.Errorf("requests异常[%v]\n", err)
				return "", err
			}
			err = Redis.Setex(fmt.Sprintf("http://%s/auth/user_token", ChatApiUrl), GetTokenResp.Data.Token, GetTokenResp.Data.ExpireTimeSeconds)
			if err != nil {
				logx.Errorf("Redis设置缓存异常[%v]\n", err)
			}
		}
	} else {
		err := requests.URL(fmt.Sprintf("http://%s/auth/user_token", ChatApiUrl)).BodyJSON(GetTokenReq).Header("operationID", "1").ToJSON(&GetTokenResp).Fetch(context.Background())
		if err != nil {
			logx.Errorf("requests异常[%v]\n", err)
			return "", err
		}
	}
	return GetTokenResp.Data.Token, nil
}
