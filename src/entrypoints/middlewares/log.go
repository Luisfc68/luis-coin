package middlewares

import (
	"github.com/google/uuid"
	"github.com/luisfc68/luis-coin/src/configuration"
	"log"
	"net/http"
)

type LogResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func newLogResponseWriter(w http.ResponseWriter) *LogResponseWriter {
	return &LogResponseWriter{ResponseWriter: w}
}

func (w *LogResponseWriter) WriteHeader(code int) {
	w.statusCode = code
	w.ResponseWriter.WriteHeader(code)
}

func LogMiddleware(server configuration.Server) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			requestId := uuid.NewString()
			logWriter := newLogResponseWriter(writer)
			log.Printf("Received request (%s) to %s (%s)\n", requestId, request.URL.Path, request.Method)
			next.ServeHTTP(logWriter, request)
			log.Printf("Sent response (%s) with status %d", requestId, logWriter.statusCode)
		})
	}
}
