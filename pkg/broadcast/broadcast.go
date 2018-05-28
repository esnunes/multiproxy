package broadcast

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"sync"
)

// Handler ...
type Handler struct {
	Addrs []*url.URL
}

func (h *Handler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)

		return
	}
	defer r.Body.Close()

	wg := sync.WaitGroup{}

	for _, addr := range h.Addrs {
		u, _ := url.Parse(r.URL.String())
		u.Scheme = addr.Scheme
		u.Host = addr.Host
		u.Path = singleJoiningSlash(addr.Path, u.Path)

		buf := bytes.NewBuffer(body)
		req, err := http.NewRequest(r.Method, u.String(), buf)
		if err != nil {
			log.Printf("Failed to create request: %v", err)
			continue
		}
		req.Header = cloneHeader(r.Header)
		if _, ok := req.Header["User-Agent"]; !ok {
			// explicitly disable User-Agent so it's not set to default value
			req.Header.Set("User-Agent", "")
		}

		wg.Add(1)
		go func() {
			defer wg.Done()

			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				log.Printf("Failed to execute request to %v: %v", u, err)
				return
			}

			if resp.StatusCode < 200 || resp.StatusCode >= 400 {
				log.Printf("Unexpected status code: %v", resp.StatusCode)
			}
		}()
	}

	wg.Wait()
}

// singleJoiningSlash merges two URL paths into one. The source code is based
// on https://golang.org/src/net/http/httputil/reverseproxy.go
func singleJoiningSlash(a, b string) string {
	aslash := strings.HasSuffix(a, "/")
	bslash := strings.HasPrefix(b, "/")
	switch {
	case aslash && bslash:
		return a + b[1:]
	case !aslash && !bslash:
		return a + "/" + b
	}
	return a + b
}

// cloneHeader creates a new Header copying all values from a base one. The
// source code is based on
// https://golang.org/src/net/http/httputil/reverseproxy.go
func cloneHeader(h http.Header) http.Header {
	h2 := make(http.Header, len(h))
	for k, vv := range h {
		vv2 := make([]string, len(vv))
		copy(vv2, vv)
		h2[k] = vv2
	}
	return h2
}
