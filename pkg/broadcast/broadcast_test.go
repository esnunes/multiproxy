package broadcast

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func TestHandler(t *testing.T) {
	type exp struct {
		Method string
		URL    *url.URL
		Header http.Header
		Body   string
		Err    error
	}

	cases := []struct {
		Name            string
		ServerCount     int
		ServerURLSuffix string
		URL             *url.URL
		Header          http.Header
		Exp             exp
	}{
		{
			Name:            "no-servers",
			ServerCount:     0,
			ServerURLSuffix: "/",
			URL:             mustParseURL("http://localhost:666/"),
			Header:          http.Header{},
			Exp: exp{
				Method: http.MethodGet,
				URL:    nil,
				Header: nil,
				Body:   "",
				Err:    nil,
			},
		},
		{
			Name:            "get",
			ServerCount:     3,
			ServerURLSuffix: "/",
			URL:             mustParseURL("http://localhost:666/"),
			Header:          http.Header{},
			Exp: exp{
				Method: http.MethodGet,
				URL:    mustParseURL("http://xxx/"),
				Header: http.Header{
					"Accept-Encoding": []string{"gzip"},
				},
				Body: "",
				Err:  nil,
			},
		},
		{
			Name:            "get:query-params",
			ServerCount:     3,
			ServerURLSuffix: "/",
			URL:             mustParseURL("http://localhost:666/?q=search"),
			Header:          http.Header{},
			Exp: exp{
				Method: http.MethodGet,
				URL:    mustParseURL("http://xxx/?q=search"),
				Header: http.Header{
					"Accept-Encoding": []string{"gzip"},
				},
				Body: "",
				Err:  nil,
			},
		},
		{
			Name:            "post",
			ServerCount:     3,
			ServerURLSuffix: "/",
			URL:             mustParseURL("http://localhost:666/level-one"),
			Header:          http.Header{},
			Exp: exp{
				Method: http.MethodPost,
				URL:    mustParseURL("http://xxx/level-one"),
				Header: http.Header{
					"Content-Length":  []string{strconv.Itoa(len("hello world"))},
					"Accept-Encoding": []string{"gzip"},
				},
				Body: "hello world",
				Err:  nil,
			},
		},
		{
			Name:            "post:suffix",
			ServerCount:     3,
			ServerURLSuffix: "/server-base-path",
			URL:             mustParseURL("http://localhost:666/level-one"),
			Header:          http.Header{},
			Exp: exp{
				Method: http.MethodPost,
				URL:    mustParseURL("http://xxx/server-base-path/level-one"),
				Header: http.Header{
					"Content-Length":  []string{strconv.Itoa(len("hello world"))},
					"Accept-Encoding": []string{"gzip"},
				},
				Body: "hello world",
				Err:  nil,
			},
		},
		{
			Name:            "post:headers",
			ServerCount:     3,
			ServerURLSuffix: "/",
			URL:             mustParseURL("http://localhost:666/level-one/level-two"),
			Header: http.Header{
				"x-app-specific": []string{"x-app-specific-value"},
			},
			Exp: exp{
				Method: http.MethodPost,
				URL:    mustParseURL("http://xxx/level-one/level-two"),
				Header: http.Header{
					"Content-Length":  []string{strconv.Itoa(len("hello world"))},
					"X-App-Specific":  []string{"x-app-specific-value"},
					"Accept-Encoding": []string{"gzip"},
				},
				Body: "hello world",
				Err:  nil,
			},
		},
	}

	asserts := func(s *spyServer, e exp, t *testing.T) {
		if s.Method != e.Method {
			t.Errorf("Expected method %v; got %v", e.Method, s.Method)
		}
		if s.URL.Path != e.URL.Path {
			t.Errorf("Expected url path %v; got %v", e.URL.Path, s.URL.Path)
		}
		if s.URL.RawQuery != e.URL.RawQuery {
			t.Errorf("Expected query %v; got %v", e.URL.RawQuery, s.URL.RawQuery)
		}
		if !reflect.DeepEqual(s.Header, e.Header) {
			t.Errorf("Expected headers %v; got %v", e.Header, s.Header)
		}
		if s.Body != e.Body {
			t.Errorf("Expected body %v; got %v", e.Body, s.Body)
		}
		if s.Err != e.Err {
			t.Errorf("Expected error %v; got %v", e.Err, s.Err)
		}
	}

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			addrs := make([]*url.URL, tc.ServerCount)
			spies := make([]*spyServer, tc.ServerCount)

			// prepare spy servers
			for i := 0; i < tc.ServerCount; i++ {
				spies[i] = &spyServer{}

				hs := httptest.NewServer(spies[i])
				defer hs.Close()

				addrs[i], _ = url.Parse(hs.URL + tc.ServerURLSuffix)
			}

			// create handler
			bh := &Handler{
				Addrs: addrs,
			}

			// prepare request
			r := httptest.NewRequest(tc.Exp.Method, tc.URL.String(), strings.NewReader(tc.Exp.Body))
			r.Header = tc.Header
			w := httptest.NewRecorder()

			bh.ServeHTTP(w, r)

			// asserts
			for _, s := range spies {
				asserts(s, tc.Exp, t)
			}
		})
	}
}

func mustParseURL(addr string) *url.URL {
	u, err := url.Parse(addr)
	if err != nil {
		panic(err)
	}
	return u
}

// spyServer stores a set of request data to be later used in assertions.
type spyServer struct {
	Method string
	URL    *url.URL
	Header http.Header
	Body   string
	Err    error
}

func (s *spyServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Method = r.Method
	s.URL = r.URL
	s.Header = r.Header

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		s.Err = err
		return
	}

	s.Body = string(body)
}
