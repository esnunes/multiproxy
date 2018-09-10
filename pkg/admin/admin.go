package admin

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/esnunes/multiproxy/pkg/static"
)

// Options ...
type Options struct {
	Debug  bool
	Config string
}

type handler struct {
	opts  Options
	index []byte
	tmpl  *template.Template
}

// NewHandler ...
func NewHandler(opts Options) http.Handler {
	t, err := static.ReadFile("admin/index.html")
	if err != nil {
		log.Fatal(err)
	}

	i := strings.Replace(
		string(t),
		"<script type=\"text/javascript\"></script>",
		fmt.Sprintf("<script type=\"text/javascript\">app.init({ debug: %v, config: %s });</script>", opts.Debug, opts.Config),
		1,
	)

	return &handler{
		index: []byte(i),
	}
}

func (h handler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		rw.Write(h.index)
	default:
		r.URL.Path = "/admin" + r.URL.Path
		static.Handler.ServeHTTP(rw, r)
	}
}
