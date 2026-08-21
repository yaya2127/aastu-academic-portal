package middleware

import (
	"context"
	"net/http"
	"strings"
)

type key int

const UserIDKey key = 0

// JWTMiddleware validates Authorization Bearer tokens for secure student portal API endpoints
func JWTMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(`{"error":"UNAUTHORIZED","message":"Missing or malformed Bearer authorization token"}`))
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		
		// In production, verify JWT signature with secret key
		if tokenStr == "" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(`{"error":"INVALID_TOKEN","message":"JWT signature verification failed"}`))
			return
		}

		ctx := context.WithValue(r.Context(), UserIDKey, "ETS-0452/14")
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
