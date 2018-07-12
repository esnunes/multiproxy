package unicast

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

// Selector is an interface that wraps the Selected method.
//
// Selected returns the key of the selected environment or nil if not found.
type Selector interface {
	Selected(*http.Request) *string
}

// Handler ...
type Handler struct {
	Selector Selector

	// Rules specifies a map between environment keys and upstreams.
	Rules map[string]*url.URL
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	v := h.Selector.Selected(r)
	if v == nil {
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	u, ok := h.Rules[*v]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	rp := httputil.NewSingleHostReverseProxy(u)
	rp.ServeHTTP(w, r)
}
