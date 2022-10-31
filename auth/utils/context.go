package utils

import (
	"context"
)

type authCtxKey string

const authorizationKey authCtxKey = "authorization"

func WithAuthorization(ctx context.Context, authorization string) context.Context {
	return context.WithValue(ctx, authorizationKey, authorization)
}

func AuthorizationFromContext(ctx context.Context) string {
	authorization, _ := ctx.Value(authorizationKey).(string)
	return authorization
}
