package middlewares

import (
	"net/http"
	"strconv"
	"time"

	"github.com/deepakvbansode/idp-cloudgenie-backend/internal/core/ports"
)

type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (r *statusRecorder) WriteHeader(code int) {
	r.status = code
	r.ResponseWriter.WriteHeader(code)
}

func InstrumentationMiddleware(logger ports.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		 return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			  start := time.Now()
			  recorder := &statusRecorder{ResponseWriter: w, status: 200}
			  next.ServeHTTP(recorder, r)
			  duration := time.Since(start)
			  logger.WithField("method", r.Method).WithField("path", r.URL.Path).WithField("status", strconv.Itoa(recorder.status)).WithField("duration_ms", strconv.FormatInt(duration.Milliseconds(), 10)).Info( "request completed")
		 })
	}
}