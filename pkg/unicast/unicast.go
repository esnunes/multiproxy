package unicast

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

// Handler ...
type Handler struct {
	Cookie   string
	Rules    map[string]*url.URL
	ErrorLog *log.Logger
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie(h.Cookie)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		h.logf("unicast: expected cookie not available: %v", h.Cookie)

		return
	}

	u, ok := h.Rules[c.Value]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)

		h.logf("unicast: expected upstream not available: %v", c.Value)

		return
	}

	rp := httputil.NewSingleHostReverseProxy(u)
	rp.ServeHTTP(w, r)
}

func (h *Handler) logf(format string, args ...interface{}) {
	if h.ErrorLog != nil {
		h.ErrorLog.Printf(format, args...)
	} else {
		log.Printf(format, args...)
	}
}
