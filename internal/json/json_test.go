package json

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRead(t *testing.T) {
	type payload struct {
		Name string `json:"name"`
	}

	cases := []struct {
		name    string
		body    string
		wantErr bool
		want    payload
	}{
		{
			name:    "valid json",
			body:    `{"name":"bryce"}`,
			wantErr: false,
			want:    payload{Name: "bryce"},
		},
		{
			name:    "malformed json",
			body:    `{"name":`,
			wantErr: true,
		},
		{
			name:    "unknown field rejected",
			body:    `{"name":"bryce","extra":"field"}`,
			wantErr: true,
		},
		{
			name:    "empty body",
			body:    ``,
			wantErr: true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(c.body))
			var got payload
			err := Read(req, &got)

			if c.wantErr && err == nil {
				t.Fatalf("expected error, got nil")
			}
			if !c.wantErr {
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
				if got != c.want {
					t.Errorf("got %+v, want %+v", got, c.want)
				}
			}
		})
	}
}

func TestWrite(t *testing.T) {
	rec := httptest.NewRecorder()
	Write(rec, http.StatusCreated, map[string]string{"id": "123"})

	if rec.Code != http.StatusCreated {
		t.Errorf("got status %d, want %d", rec.Code, http.StatusCreated)
	}
	if ct := rec.Header().Get("Content-Type"); ct != "application/json" {
		t.Errorf("got Content-Type %q, want %q", ct, "application/json")
	}

	var body map[string]string
	if err := json.NewDecoder(rec.Body).Decode(&body); err != nil {
		t.Fatalf("failed to decode response body: %v", err)
	}
	if body["id"] != "123" {
		t.Errorf("got id %q, want %q", body["id"], "123")
	}
}

func TestWriteError(t *testing.T) {
	rec := httptest.NewRecorder()
	WriteError(rec, http.StatusBadRequest, "invalid input")

	if rec.Code != http.StatusBadRequest {
		t.Errorf("got status %d, want %d", rec.Code, http.StatusBadRequest)
	}

	var body errorResponse
	if err := json.NewDecoder(rec.Body).Decode(&body); err != nil {
		t.Fatalf("failed to decode response body: %v", err)
	}
	if body.Error != "invalid input" {
		t.Errorf("got error %q, want %q", body.Error, "invalid input")
	}
}
