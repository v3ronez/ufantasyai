package view

import (
	"context"
	"strconv"

	"github.com/v3ronez/ufantasyai/types"
)

func AuthenticatedUser(ctx context.Context) types.AuthenticateUser {
	user, ok := ctx.Value(types.UserContextKey).(types.AuthenticateUser)
	if !ok {
		return types.AuthenticateUser{}
	}
	return user
}

func String(n int) string {
	return strconv.Itoa(n)
}
