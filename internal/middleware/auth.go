package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/Alfeus22/puskesmas-api/pkg/utils"
)

// contextKey digunakan sebagai kunci rahasia agar tidak bentrok dengan package lain
type contextKey struct {
	name string
}

var userCtxKey = &contextKey{"user"}

// AuthMiddleware
func AuthMiddleware() func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Authorization")

			if header == "" {
				h.ServeHTTP(w, r)
				return
			}
			bearerToken := strings.Split(header, " ")
			if len(bearerToken) != 2 || bearerToken[0] != "Bearer" {
				http.Error(w, "Format token tidak valid", http.StatusUnauthorized)
				return
			}

			tokenString := bearerToken[1]

			// Minta utils kita untuk memeriksa keaslian token
			claims, err := utils.ValidatedToken(tokenString)
			if err != nil {
				http.Error(w, "Token tidak valid atau kedaluwarsa", http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), userCtxKey, claims)
			req := r.WithContext(ctx)
			h.ServeHTTP(w, req)
		})
	}
}

func ForContext(ctx context.Context) *utils.JwtCustomClaims {
	raw, _ := ctx.Value(userCtxKey).(*utils.JwtCustomClaims)
	return raw
}
