package helpers

import (
	"context"
)

type (
	Key string
)

const (
	ContextKeyToken = Key("token_context")
	ContextKeyAuth  = Key("auth_userID")
)

// SetTokenContext set the token string data to the context
func SetTokenContext(parent context.Context, value string) context.Context {
	if parent == nil {
		parent = context.Background()
	}

	parent = context.WithValue(parent, ContextKeyToken, value)
	return parent
}

// GetTokenContext get the token string data from the context
func GetTokenContext(ctx context.Context) (token string, ok bool) {
	if ctx == nil {
		return
	}
	if val := ctx.Value(ContextKeyToken); val != nil {
		token, ok = val.(string)
	}
	return
}

// SetUserIDContext set the userID data to the context
func SetUserIDContext(parent context.Context, value int64) context.Context {
	if parent == nil {
		parent = context.Background()
	}

	parent = context.WithValue(parent, ContextKeyAuth, value)
	return parent
}

// GetUserIDContext get the userID data from the context
func GetUserIDContext(ctx context.Context) (userID int64, ok bool) {
	if ctx == nil {
		return
	}
	if val := ctx.Value(ContextKeyToken); val != nil {
		userID, ok = val.(int64)
	}
	return
}
