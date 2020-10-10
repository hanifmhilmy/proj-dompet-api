package helpers

import (
	"context"
)

type (
	Key string
)

const (
	ContextKeyToken    = Key("token_context")
	ContextKeyAuth     = Key("auth_userID")
	ContextKeyAuthUUID = Key("auth_userUUID")
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
	if val := ctx.Value(ContextKeyAuth); val != nil {
		userID, ok = val.(int64)
	}
	return
}

// SetUserUUIDContext set the userID data to the context
func SetUserUUIDContext(parent context.Context, value string) context.Context {
	if parent == nil {
		parent = context.Background()
	}

	parent = context.WithValue(parent, ContextKeyAuthUUID, value)
	return parent
}

// GetUserUUIDContext get the userID data from the context
func GetUserUUIDContext(ctx context.Context) (UUID string, ok bool) {
	if ctx == nil {
		return
	}
	if val := ctx.Value(ContextKeyAuthUUID); val != nil {
		UUID, ok = val.(string)
	}
	return
}
