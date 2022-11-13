package util

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"google.golang.org/grpc/metadata"
)

func GetTokenFromCtx(ctx context.Context) (token string, err error) {
	metas, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		err = fmt.Errorf("%+v", metas)
		return
	}
	tokens := metas.Get("authorization")
	if tokens == nil || len(tokens) == 0 {
		err = fmt.Errorf("invalid token")
		return
	}
	token = tokens[0]
	return
}

func GenerateToken() (string, string) {
	token := uuid.NewString()
	refreshToken := uuid.NewString()
	return token, refreshToken
}
