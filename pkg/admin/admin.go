package admin

import (
	"html/template"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/esnunes/multiproxy"
	"github.com/esnunes/multiproxy/pkg/static"
)

// Options ...
type Options struct {
	Cookie    string
	Upstreams []multiproxy.Upstream
	Broadcast []string
}

type handler struct {
	opts Options
	tmpl *template.Template
}

// NewHandler ...
func NewHandler(opts Options) http.Handler {
	tmplb, err := static.ReadFile("admin/index.html")
	if err != nil {
		log.Fatal(err)
	}

	tmpl := template.Must(template.New("index").Parse(string(tmplb)))

	h := &handler{
		opts: opts,
		tmpl: tmpl,
	}

	return h
}

func (h handler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.show(rw, r)
	case http.MethodPost:
		h.save(rw, r)
	}
}

func (h handler) show(rw http.ResponseWriter, r *http.Request) {
	selected := ""
	status := ""

	if v, ok := r.URL.Query()["status"]; ok {
		status = strings.Join(v, "")
	}

	if c, err := r.Cookie(h.opts.Cookie); err == nil {
		selected = c.Value
	}

	data := struct {
		Options  Options
		Selected string
		Status   string
	}{
		Options:  h.opts,
		Selected: selected,
		Status:   status,
	}

	h.tmpl.Execute(rw, data)
}

func (h handler) save(rw http.ResponseWriter, r *http.Request) {
	status := "error"

	r.ParseForm()

	if v, ok := r.Form["upstream"]; ok {
		value := strings.Join(v, "")

		http.SetCookie(rw, &http.Cookie{
			Name:    h.opts.Cookie,
			Value:   value,
			Expires: time.Now().Add(365 * 24 * time.Hour),
		})

		status = "ok"
	}

	http.Redirect(rw, r, r.URL.Path+"?status="+status, http.StatusSeeOther)
}
