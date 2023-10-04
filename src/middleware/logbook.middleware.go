package myMiddleware

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	_logs "github.com/FranMT-S/chi-zinc-server/src/logs"
)

/*
LogBookMiddleware creates a log to keep track of the routes consulted and values sent to the server
*/
func LogBookMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		body := ""
		buffer := new(bytes.Buffer)
		bytes, _ := io.ReadAll(r.Body)
		json.Compact(buffer, bytes)

		body = buffer.String()
		_logs.LogBookSVG(r.Method, r.URL.Path, body)

		next.ServeHTTP(w, r)
	})
}
