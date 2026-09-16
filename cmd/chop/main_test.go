package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMount(t *testing.T) {
	app := &application{}
	ts := httptest.NewServer(app.mount())
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/v1/healthz")
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("TestMount (/v1/healthz): got %d, want 200", resp.StatusCode)
	}
}
