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

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://site.example.com:8080/", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler.ServeHTTP(recorder, req)

	t.Helper()

	if req.URL.Path != "/site.example.com/index.html" {
		t.Errorf("invalid path value: %s", req.URL.Path)
	}
}

func TestDemoSubfolder(t *testing.T) {
	cfg := traefik_plugin_static_sites.CreateConfig()

	ctx := context.Background()
	next := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {})

	handler, err := traefik_plugin_static_sites.New(ctx, next, cfg, "staticsites-plugin")
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://site.example.com/dir/", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler.ServeHTTP(recorder, req)

	t.Helper()

	if req.URL.Path != "/site.example.com/dir/index.html" {
		t.Errorf("invalid path value: %s", req.URL.Path)
	}
}