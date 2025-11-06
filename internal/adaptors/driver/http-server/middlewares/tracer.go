package middlewares

import (
	"context"
	"net/http"

	"github.com/deepakvbansode/idp-cloudgenie-backend/internal/common/constants"
	"github.com/google/uuid"
)

type contextKey string

func TracerMiddleware() func(http.Handler) http.Handler {
       return func(next http.Handler) http.Handler {
	       return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		       traceId := uuid.New().String()
		       ctx := context.WithValue(r.Context(), constants.TraceIDKey, traceId)
		       next.ServeHTTP(w, r.WithContext(ctx))
	       })
       }
}
