package env

import (
	"errors"
	"os"
	"testing"
)

func TestGetString(t *testing.T) {
	cases := []struct {
		name     string
		key      string
		envVal   string
		setEnv   bool
		fallback string
		want     string
	}{
		{
			name:     "returns env value when set",
			key:      "TEST_GETSTRING_SET",
			envVal:   "hello",
			setEnv:   true,
			fallback: "fallback",
			want:     "hello",
		},
		{
			name:     "returns fallback when not set",
			key:      "TEST_GETSTRING_UNSET",
			setEnv:   false,
			fallback: "fallback",
			want:     "fallback",
		},
		{
			name:     "returns empty string when set to empty",
			key:      "TEST_GETSTRING_EMPTY",
			envVal:   "",
			setEnv:   true,
			fallback: "fallback",
			want:     "",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if c.setEnv {
				os.Setenv(c.key, c.envVal)
				defer os.Unsetenv(c.key)
			} else {
				os.Unsetenv(c.key)
			}

			got := GetString(c.key, c.fallback)
			if got != c.want {
				t.Errorf("GetString(%q, %q) = %q, want %q", c.key, c.fallback, got, c.want)
			}
		})
	}
}

func TestGetRequiredString(t *testing.T) {
	cases := []struct {
		name    string
		key     string
		envVal  string
		setEnv  bool
		want    string
		wantErr error
	}{
		{
			name:   "returns value when set",
			key:    "TEST_REQUIRED_SET",
			envVal: "abc123",
			setEnv: true,
			want:   "abc123",
		},
		{
			name:    "returns error when not set",
			key:     "TEST_REQUIRED_UNSET",
			setEnv:  false,
			want:    "",
			wantErr: ErrEnvVarNotFound,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if c.setEnv {
				os.Setenv(c.key, c.envVal)
				defer os.Unsetenv(c.key)
			} else {
				os.Unsetenv(c.key)
			}

			got, err := GetRequiredString(c.key)
			if got != c.want {
				t.Errorf("GetRequiredString(%q) = %q, want %q", c.key, got, c.want)
			}
			if !errors.Is(err, c.wantErr) {
				t.Errorf("GetRequiredString(%q) = err %v, want %v", c.key, err, c.wantErr)
			}
		})
	}
}
