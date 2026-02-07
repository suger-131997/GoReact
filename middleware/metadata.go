package middleware

import (
	"goreact/internal/contextutil"
	"net/http"
)

func MetadataMiddleware(next http.Handler, defaultTitle string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		ctx = contextutil.WithTitle(ctx, defaultTitle)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
