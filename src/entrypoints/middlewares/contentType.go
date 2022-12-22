package middlewares

import (
	"github.com/luisfc68/luis-coin/src/configuration"
	"net/http"
)

func ContentTypeMiddleware(server configuration.Server) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			writer.Header().Set("Content-Type", "application/json")
			next.ServeHTTP(writer, request)
		})
	}
}
