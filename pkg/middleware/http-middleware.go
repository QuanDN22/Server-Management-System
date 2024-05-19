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
		if r.URL.Path == "/v1/api/login" || r.URL.Path == "/v1/api/signup" {
			fmt.Println("2. start /login")
			// check basic auth
			// _, _, ok := r.BasicAuth()
			// if !ok {
			// 	w.WriteHeader(http.StatusUnauthorized)
			// 	w.Write([]byte("missing basic auth")) //nolint
			// 	return
			// }

			// fmt.Println("http 1")
			// var v map[string]string
			// err := json.NewDecoder(r.Body).Decode(&v)
			// if err != nil {
			// 	w.WriteHeader(http.StatusBadRequest)
			// 	w.Write([]byte("invalid json body")) //nolint
			// 	return
			// }

			// // fmt.Println("http 2")

			// fmt.Println(v["username"], v["password"])
			// if v["username"] == "" || v["password"] == "" {
			// 	w.WriteHeader(http.StatusUnauthorized)
			// 	w.Write([]byte("invalid credentials, missing username or password")) //nolint
			// 	return
			// }

			// fmt.Println("http 3")

			// // // r.SetBasicAuth(v["username"], v["password"])
			// r.Header.Add("Content-Type", "application/json")
			// json.NewDecoder(r.Body).Decode(&v)
			// fmt.Println(v["username"], v["password"])

			// fmt.Println("http 4")

			next.ServeHTTP(w, r)

			// fmt.Println("http 5")

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
