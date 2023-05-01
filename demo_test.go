package traefik_plugin_static_sites_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/russ-p/traefik-plugin-static-sites"
)

func TestDemo(t *testing.T) {
	cfg := traefik_plugin_static_sites.CreateConfig()

	ctx := context.Background()
	next := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {})

	handler, err := traefik_plugin_static_sites.New(ctx, next, cfg, "staticsites-plugin")
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://site.example.com/", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler.ServeHTTP(recorder, req)

	t.Helper()

	if req.URL.Path == "http://site.example.com/site.example.com/index.html" {
		t.Errorf("invalid path value: %s", req.URL.Path)
	}
}