package middleware

import (
	"fmt"
	"net/http"
	"strings"
)

type Middleware struct {
	Validator
}

func NewMiddleware(publicKeyPath string) (*Middleware, error) {
	validator, err := NewValidator(publicKeyPath)
	if err != nil {
		return nil, fmt.Errorf("unable to create validator: %w", err)
	}

	return &Middleware{
		Validator: *validator,
	}, nil
}

func (mw *Middleware) HandleHTTP(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/v1/api/auth/login" || r.URL.Path == "/v1/api/auth/signup" || r.URL.Path == "/v1/api/auth/hello" {
			fmt.Println("2. start /login")
			next.ServeHTTP(w, r)

			fmt.Println("2. finish /login")
			return
		}

		fmt.Println("2. start")

		parts := strings.Split(r.Header.Get("Authorization"), " ")
		if len(parts) < 2 || parts[0] != "Bearer" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("missing or invalid authorization header")) //nolint
			return
		}
		tokenString := parts[1]

		token, err := mw.GetToken(tokenString)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("invalid token: " + err.Error())) //nolint
			return
		}

		ctx := ContextSetToken(r.Context(), token)

		// call the next handler with the updated context
		next.ServeHTTP(w, r.WithContext(ctx))

		fmt.Println("2. finish")
	}
}
