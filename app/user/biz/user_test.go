package biz

import (
	user "eshop/api/user/v1"
	"eshop/app/user/internal/model"
	"eshop/app/user/internal/util"
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBusinessHandler_CreateUser(t *testing.T) {
	user := &model.User{
		Id:     util.GenerateID(),
		Email:  "test1@email.com",
		Nick:   "test1",
		Passwd: "111111",
		Status: 1,
		Mobile: "13798296811",
	}
	err := bizHandler.CreateUser(user)
	require.NoError(t, err)
}

func TestBusinessHandler_GetUser(t *testing.T) {
	id := int64(12434488763879424)
	user, err := bizHandler.GetUser(id)
	require.NoError(t, err)
	fmt.Printf("%+v\n", user)
}

func TestBusinessHandler_UpdateUser(t *testing.T) {
	user := &model.User{
		Id:     12434488763879424,
		Email:  "update@email.com",
		Nick:   "update",
		Passwd: "111111",
		Status: 1,
		Mobile: "13798296811",
	}
	err := bizHandler.UpdateUser(user)
	require.NoError(t, err)
}

func TestBusinessHandler_DeleteUser(t *testing.T) {
	id := int64(12434488763879424)
	err := bizHandler.DeleteUser(id)
	require.NoError(t, err)
}

func TestBusinessHandler_ListUser(t *testing.T) {
	filter := &user.ListFilter{
		Name:  "",
		Nick:  "update",
		Email: "",
		Page: &user.Page{
			Page:  0,
			Limit: 1,
		},
	}
	list, total, err := bizHandler.ListUser(filter)
	require.NoError(t, err)
	for _, v := range list {
		fmt.Println(v)
	}
	fmt.Println(total)
}
