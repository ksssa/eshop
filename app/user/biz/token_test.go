package biz

import (
	"eshop/app/user/internal/model"
	"eshop/app/user/internal/util"
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestBusinessHandler_CheckToken(t *testing.T) {
	token := "907ab9b8-3a4e-499a-8bc4-f0f0265c970e"
	userId, isAdmin, err := bizHandler.CheckToken(token)
	require.NoError(t, err)
	fmt.Println(userId, isAdmin)
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
	token := "907ab9b8-3a4e-499a-8bc4-f0f0265c970e"
	err := bizHandler.DeleteToken(token)
	require.NoError(t, err)
}
