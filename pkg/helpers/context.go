package helpers

import (
	"context"
)

type (
	Key string
)

const (
	ContextKeyToken = Key("token_context")
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
