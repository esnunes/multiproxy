package envs

import (
	"encoding/json"
	"net/http"
	"time"
)

// Handler ...
type Handler struct {
	// Cookie specifies the name of the cookie used to store and retrieve
	// selected environment.
	Cookie string
}

// Selected returns the key of the selected environment or nil if not found.
func (h *Handler) Selected(r *http.Request) *string {
	return getCookieValue(r, h.Cookie)
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		data := body{
			Key: h.Selected(r),
		}

		json.NewEncoder(w).Encode(data)
	case http.MethodPost:
		var data body

		defer r.Body.Close()
		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			w.WriteHeader(http.StatusBadRequest)

			return
		}

		setCookieValue(w, h.Cookie, data.Key)

		json.NewEncoder(w).Encode(data)
	}
}

type body struct {
	Key *string `json:"key"`
}

// getCookieValue returns the value of the named cookie provided in the request
// or nil if not found.
func getCookieValue(r *http.Request, name string) *string {
	c, err := r.Cookie(name)
	if err != nil {
		return nil
	}

	return &c.Value
}

// setCookieValue stores the given value in the named cookie or expires cookie
// if nil.
func setCookieValue(w http.ResponseWriter, name string, v *string) {
	c := &http.Cookie{
		Name:    name,
		Expires: time.Unix(0, 0),
	}

	if v != nil {
		c.Value = *v
		c.Expires = time.Now().Add(356 * 24 * time.Hour)
	}

	http.SetCookie(w, c)
}
