package cors

import (
	"net/http"
)

// Cors ...
type Cors struct {
	Origin string
}

// Handler ...
func (c *Cors) Handler(h http.Handler) http.HandlerFunc {
	if c.Origin == "" {
		return h.ServeHTTP
	}

	return func(w http.ResponseWriter, r *http.Request) {
		headers := w.Header()
		headers.Set("Access-Control-Allow-Origin", c.Origin)
		headers.Set("Access-Control-Allow-Credentials", "true")

		if r.Method == http.MethodOptions {
			headers.Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
			headers.Set("Access-Control-Allow-Headers", "Content-Type, *")
			headers.Set("Access-Control-Max-Age", "86400")

			headers.Add("Vary", "Origin")
			headers.Add("Vary", "Access-Control-Request-Method")
			headers.Add("Vary", "Access-Control-Request-Headers")

			return
		}

		h.ServeHTTP(w, r)
	}
}
