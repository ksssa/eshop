package biz

import (
	"eshop/app/user/internal/model"
	"eshop/app/user/internal/util"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestBusinessHandler_CheckToken(t *testing.T) {

}

func TestBusinessHandler_CreateToken(t *testing.T) {
	accessToken, refreshToken := util.GenerateToken()
	token := &model.Token{
		Id:                 util.GenerateID(),
		IsAdmin:            false,
		UserId:             12434488763879424,
		AccessToken:        accessToken,
		AccessTokenExpire:  time.Now().Unix() + 24*60*60,
		RefreshToken:       refreshToken,
		RefreshTokenExpire: time.Now().Unix() + 30*24*60*60,
	}
	err := bizHandler.CreateToken(token)
	require.NoError(t, err)
}

func TestBusinessHandler_DeleteToken(t *testing.T) {
	token := ""
	err := bizHandler.DeleteToken(token)
	require.NoError(t, err)
}
