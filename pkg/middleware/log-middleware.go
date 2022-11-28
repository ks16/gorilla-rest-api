package middleware

import (
	"context"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"

	"github.com/ks16/gorilla-rest-api/pkg/log"
)

const (
	CorrelationIdKey      = "correlation_id"
	CorrelationHeaderName = "X-Request-ID"
)

type ContextKey string

func LoggerMiddleware(logger *log.Logger) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			id := req.Header.Get(CorrelationHeaderName)
			if id == "" {
				id = uuid.New().String()
			}

			correlationCtx := context.WithValue(req.Context(), ContextKey(CorrelationIdKey), id)
			lg := logger.With().Str(CorrelationIdKey, id).Logger()
			ctxLog := lg.WithContext(correlationCtx)
			r := req.WithContext(ctxLog)

			next.ServeHTTP(w, r)
		})
	}
}
