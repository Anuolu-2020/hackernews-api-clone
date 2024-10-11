package auth

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/Anuolu-2020/hackernews-api-clone/internal/users"
	"github.com/Anuolu-2020/hackernews-api-clone/pkg/token"
)

type contextKey struct {
	name string
}

var userCtxKey = &contextKey{"user"}

func Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Authorization")

			if header == "" {
				next.ServeHTTP(w, r)
				return
			}

			// Validate token
			tokenStr := header
			username, err := token.ParseToken(tokenStr)
			if err != nil {
				log.Printf("Token ERROR: %v", err)
				http.Error(w, "Invalid Token", http.StatusForbidden)
			}

			// Check user in db
			user := users.User{Username: username}
			id, err := users.GetUserIdByUsername(username)
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}

			user.ID = strconv.Itoa(id)

			// Pass user object to request context
			ctx := context.WithValue(r.Context(), userCtxKey, &user)

			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

// finds the user from the context.
func ForContext(ctx context.Context) *users.User {
	raw, _ := ctx.Value(userCtxKey).(*users.User)

	return raw
}
