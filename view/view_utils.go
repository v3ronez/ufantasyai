package view

import (
	"context"

	"github.com/v3ronez/ufantasyai/types"
)

func AuthenticatedUser(ctx context.Context) types.AuthenticateUser {
	user, ok := ctx.Value(types.UserContextKey).(types.AuthenticateUser)
	if !ok {
		return types.AuthenticateUser{}
	}
	return user
}
