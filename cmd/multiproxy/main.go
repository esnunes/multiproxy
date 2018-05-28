package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/esnunes/multiproxy/pkg/broadcast"
	"github.com/esnunes/multiproxy/pkg/unicast"
)

// Upstream ...
type Upstream struct {
	Key  string `json:"key"`
	Addr string `json:"addr"`
}

// Config ...
type Config struct {
	Admin     string     `json:"admin"`
	Cookie    string     `json:"cookie"`
	Upstreams []Upstream `json:"upstreams"`
	Broadcast []string   `json:"broadcast"`
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
		c.Admin = "/_multiproxy"
	}
	if c.Cookie == "" {
		c.Cookie = "multiproxy"
	}
	if c.Upstreams == nil {
		c.Upstreams = []Upstream{}
	}
	if c.Broadcast == nil {
		c.Broadcast = []string{}
	}

	return &c, nil
}

// ParseUpstreams ...
func ParseUpstreams(us []Upstream) ([]*url.URL, map[string]*url.URL, error) {
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
	mux.HandleFunc(c.Admin, func(w http.ResponseWriter, r *http.Request) {
	})

	// broadcast
	bh := &broadcast.Handler{
		Addrs: addrs,
	}
	for _, endp := range c.Broadcast {
		mux.Handle(endp, bh)
	}

	// unicast
	mux.Handle("/", &unicast.Handler{
		Cookie: c.Cookie,
		Rules:  rules,
	})

	http.ListenAndServe(":8080", mux)
}
