package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/esnunes/multiproxy/pkg/admin"
	"github.com/esnunes/multiproxy/pkg/broadcast"
	"github.com/esnunes/multiproxy/pkg/cors"
	"github.com/esnunes/multiproxy/pkg/envs"
	"github.com/esnunes/multiproxy/pkg/unicast"
)

// Environment ...
type Environment struct {
	Name     string `json:"name"`
	Key      string `json:"key"`
	Upstream string `json:"upstream"`
}

// App ...
type App struct {
	ID           int           `json:"id"`
	Name         string        `json:"name"`
	Description  string        `json:"description"`
	Broadcast    []string      `json:"broadcast"`
	Addr         string        `json:"addr"`
	Environments []Environment `json:"envs"`
}

// Config ...
type Config struct {
	Admin  string `json:"admin"`
	Cookie string `json:"cookie"`
	Apps   []App  `json:"apps"`
}

// LoadConfigFromFile ...
func LoadConfigFromFile(p string) (*Config, string, error) {
	var c Config

	rawC, err := ioutil.ReadFile(p)
	if err != nil {
		return nil, "", err
	}

	if err := json.Unmarshal(rawC, &c); err != nil {
		return nil, "", err
	}

	return &c, string(rawC), nil
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

// ParseEnvironments ...
func ParseEnvironments(envs []Environment) ([]*url.URL, map[string]*url.URL, error) {
	addrs := make([]*url.URL, len(envs))
	rules := map[string]*url.URL{}

	for i, env := range envs {
		u, err := url.Parse(env.Upstream)
		if err != nil {
			return nil, nil, err
		}

		addrs[i] = u
		rules[env.Key] = u
	}

	return addrs, rules, nil
}

func main() {
	log.Printf("multiproxy: Multicast HTTP Reverse Proxy")

	if len(os.Args) < 2 {
		log.Fatal("Usage: multiproxy ./path/to/config.json")
	}

	c, rawC, err := LoadConfigFromFile(os.Args[1])
	if err != nil {
		log.Fatalf("Failed to load config file [%v]: %v", os.Args[1], err)
	}

	mux := http.NewServeMux()

	// admin
	mux.Handle(PatternFromAddr(c.Admin), admin.NewHandler(admin.Options{
		Debug:  false,
		Config: rawC,
	}))

	ch := cors.Cors{Origin: OriginFromAddr(c.Admin)}

	for _, app := range c.Apps {
		upstreams, rules, err := ParseEnvironments(app.Environments)
		if err != nil {
			log.Fatalf("Failed to parse environments [%v]: %v", app.Environments, err)
		}

		// broadcast
		bh := &broadcast.Handler{
			Addrs: upstreams,
		}
		for _, endp := range app.Broadcast {
			mux.Handle(PatternFromAddr(app.Addr)+endp, bh)
		}

		eh := &envs.Handler{Cookie: c.Cookie}
		mux.HandleFunc(PatternFromAddr(app.Addr)+"_multiproxy", ch.Handler(eh))

		// unicast
		mux.Handle(PatternFromAddr(app.Addr), &unicast.Handler{
			Selector: eh,
			Rules:    rules,
		})
	}

	log.Print("Listening at: 0.0.0.0:8080")
	http.ListenAndServe(":8080", mux)
}
