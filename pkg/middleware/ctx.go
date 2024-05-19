package middleware

import (
	"context"
	"errors"

	"github.com/golang-jwt/jwt"
)

func ContextSetToken(ctx context.Context, token *jwt.Token) context.Context {
	// return context.WithValue(ctx, tokenContextKey, token)
	return context.WithValue(ctx, "token", token)
}

func ContextGetToken(ctx context.Context) (*jwt.Token, error) {
	val := ctx.Value("token")

	if val == nil {
		return nil, errors.New("no token in context")
	}

	t, ok := val.(*jwt.Token)
	if !ok {
		return nil, errors.New("unexpected token type in context")
	}

	return t, nil
}
