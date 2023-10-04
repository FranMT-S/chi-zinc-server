package myMiddleware

import "net/http"

// JsonMiddleware set Content-Type:application/json in the header response
func JsonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		next.ServeHTTP(w, r)
	})
}
