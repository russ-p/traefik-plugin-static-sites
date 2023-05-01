// Package 
package traefik_plugin_static_sites

import (
	// "bytes"
	"context"
	"fmt"
	"net/http"
	"strings"
)

// Config the plugin configuration.
type Config struct {
	SpaFriendly bool `json:"spaFriendly"`
}

// CreateConfig creates the default plugin configuration.
func CreateConfig() *Config {
	return &Config{
		SpaFriendly: false,
	}
}

// Demo a Demo plugin.
type Demo struct {
	next     	 http.Handler
	spaFriendly  bool
	name         string
}

// New created a new Demo plugin.
func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {

	return &Demo{
		spaFriendly:  config.SpaFriendly,
		next:     next,
		name:     name,
	}, nil
}

func (a *Demo) ServeHTTP(rw http.ResponseWriter, req *http.Request) {

	host :=  strings.Split(req.Host, ":")[0]
	path := "/" + host + req.URL.Path

	// is dir?
	if strings.HasSuffix(path, "/") {
		path = path + "index.html"
	}

	fmt.Printf("Forward to %s\n", path)

	req.URL.Path = path
	a.next.ServeHTTP(rw, req)
}
