package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/esnunes/multiproxy"
	"github.com/esnunes/multiproxy/pkg/admin"
	"github.com/esnunes/multiproxy/pkg/broadcast"
	"github.com/esnunes/multiproxy/pkg/cors"
	"github.com/esnunes/multiproxy/pkg/envs"
	"github.com/esnunes/multiproxy/pkg/unicast"
)

// Config ...
type Config struct {
	Admin     string                `json:"admin"`
	Cookie    string                `json:"cookie"`
	Upstreams []multiproxy.Upstream `json:"upstreams"`
	Broadcast []string              `json:"broadcast"`
}

func (c Config) String() string {
	var b strings.Builder

	fmt.Fprintf(&b, "Admin:\t%v\n", c.Admin)
	fmt.Fprintf(&b, "Cookie:\t%v\n", c.Cookie)
	fmt.Fprintf(&b, "Upstream:\n")
	for _, up := range c.Upstreams {
		fmt.Fprintf(&b, "\t- Key: %v, Addr: %v\n", up.Key, up.Addr)
	}
	fmt.Fprintf(&b, "Broadcast:\n")
	for _, endp := range c.Broadcast {
		fmt.Fprintf(&b, "\t- %v\n", endp)
	}

	return b.String()
}

// LoadConfigFromFile ...
func LoadConfigFromFile(p string) (*Config, error) {
	var c Config

	f, err := os.Open(p)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	if err := json.NewDecoder(f).Decode(&c); err != nil {
		return nil, err
	}

	// set default values
	if c.Admin == "" {
		c.Admin = "/_admin"
	}
	if c.Cookie == "" {
		c.Cookie = "multiproxy"
	}
	if c.Upstreams == nil {
		c.Upstreams = []multiproxy.Upstream{}
	}
	if c.Broadcast == nil {
		c.Broadcast = []string{}
	}

	return &c, nil
}

// PatternFromAddr returns a ServeMux compatible pattern based on the given URL.
func PatternFromAddr(a string) string {
	u, _ := url.Parse(a)

	p := u.Path
	if !strings.HasSuffix(p, "/") {
		p = p + "/"
	}

	return u.Hostname() + p
}

// OriginFromAddr returns a Origin Header compatible value based on the given URL.
func OriginFromAddr(a string) string {
	u, _ := url.Parse(a)

	if u.Scheme == "" || u.Host == "" {
		return ""
	}

	return u.Scheme + "://" + u.Host
}

// ParseUpstreams ...
func ParseUpstreams(us []multiproxy.Upstream) ([]*url.URL, map[string]*url.URL, error) {
	addrs := make([]*url.URL, len(us))
	rules := map[string]*url.URL{}

	for i, up := range us {
		u, err := url.Parse(up.Addr)
		if err != nil {
			return nil, nil, err
		}

		addrs[i] = u
		rules[up.Key] = u
	}

	return addrs, rules, nil
}

func main() {
	log.Printf("multiproxy: Multicast HTTP Reverse Proxy")

	if len(os.Args) < 2 {
		log.Fatal("Usage: multiproxy ./path/to/config.json")
	}

	c, err := LoadConfigFromFile(os.Args[1])
	if err != nil {
		log.Fatalf("Failed to load config file [%v]: %v", os.Args[1], err)
	}
	log.Printf("Config:\n%v", c)

	addrs, rules, err := ParseUpstreams(c.Upstreams)
	if err != nil {
		log.Fatalf("Failed to parse upstreams [%v]: %v", c.Upstreams, err)
	}

	mux := http.NewServeMux()

	// admin
	mux.Handle(PatternFromAddr(c.Admin), admin.NewHandler(admin.Options{
		Cookie:    c.Cookie,
		Upstreams: c.Upstreams,
		Broadcast: c.Broadcast,
	}))

	// broadcast
	bh := &broadcast.Handler{
		Addrs: addrs,
	}
	for _, endp := range c.Broadcast {
		mux.Handle(endp, bh)
	}

	ch := cors.Cors{Origin: OriginFromAddr(c.Admin)}

	eh := &envs.Handler{Cookie: c.Cookie}
	mux.HandleFunc("/_multiproxy", ch.Handler(eh))

	// unicast
	mux.Handle("/", &unicast.Handler{
		Selector: eh,
		Rules:    rules,
	})

	log.Print("Listening at: 0.0.0.0:8080")
	http.ListenAndServe(":8080", mux)
}
